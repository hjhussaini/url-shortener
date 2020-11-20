package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/hjhussaini/url-shortener/database"
	"github.com/hjhussaini/url-shortener/keygen-server/models"
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
		fmt.Println("ParseInt", err.Error())
		os.Exit(1)
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	cassandra, err := database.CassandraConnect(databaseServer, databaseKeyspace)
	if err != nil {
		fmt.Println("Cassandra", err.Error())
		os.Exit(1)
	}
	defer cassandra.Close()

	go keyGenerator(cassandra, time.Duration(keyGenerationInterval))

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
	keys := models.Keys{}
	usedKeys := models.UsedKeys{}
	number := keys.Count(session) + usedKeys.Count(session)

	time.Sleep(interval * time.Millisecond)
	for range time.Tick(interval * time.Millisecond) {
		number = number + 1
		keys.Key = generateKey(number)
		if err := keys.Insert(session); err != nil {
			fmt.Println(err.Error())
		}
	}
}
