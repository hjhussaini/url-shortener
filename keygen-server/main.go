package main

import (
	"crypto/md5"
	"encoding/base64"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/hjhussaini/url-shortener/cache"
	"github.com/hjhussaini/url-shortener/database"
	"github.com/hjhussaini/url-shortener/keygen-server/api"
	"github.com/hjhussaini/url-shortener/keygen-server/models"
	"github.com/hjhussaini/url-shortener/logger"
	"github.com/hjhussaini/url-shortener/server"
)

func main() {
	databaseServer := os.Getenv("DATABASE_SERVER")
	databaseKeyspace := os.Getenv("DATABASE_KEYSPACE")
	keyGenerationInterval, err := strconv.ParseInt(
		os.Getenv("KEY_GENERATION_INTERVAL"),
		10,
		64,
	)
	if err != nil {
		logger.Fatal(err)
	}
	cacheServer := os.Getenv("CACHE_SERVER")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	cassandra, err := database.CassandraConnect(databaseServer, databaseKeyspace)
	if err != nil {
		logger.Fatal(err)
	}
	defer cassandra.Close()

	go keyGenerator(cassandra, time.Duration(keyGenerationInterval))

	redisCache := cache.NewRedisCache(cacheServer, 0, "keys_cache")
	api := api.API{
		Session: cassandra,
		Cache:   redisCache,
	}
	api.Caching()

	http.HandleFunc("/keys", api.GetKey)

	server.RunHTTP(host+":"+port, nil)
}

func generateKey(number int64) string {
	md5Encoded := md5.Sum([]byte(strconv.FormatInt(number, 10)))
	bytes := make([]byte, 0)
	for _, value := range md5Encoded {
		bytes = append(bytes, value)
	}
	base64Encoded := base64.StdEncoding.EncodeToString(bytes)
	key := string(base64Encoded)

	return key[1:7]
}

func keyGenerator(session database.Session, interval time.Duration) {
	keys := models.Keys{Session: session}
	usedKeys := models.UsedKeys{Session: session}
	number := keys.Count() + usedKeys.Count()

	time.Sleep(interval * time.Millisecond)
	for range time.Tick(interval * time.Millisecond) {
		number = number + 1
		keys.Key = generateKey(number)
		if err := keys.Insert(); err != nil {
			logger.Error(err)
		}
	}
}
