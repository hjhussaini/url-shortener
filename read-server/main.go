package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/hjhussaini/url-shortener/cache"
	"github.com/hjhussaini/url-shortener/database"
	"github.com/hjhussaini/url-shortener/logger"
	"github.com/hjhussaini/url-shortener/read-server/api"
	"github.com/hjhussaini/url-shortener/server"
)

func main() {
	databaseServer := os.Getenv("DATABASE_SERVER")
	databaseKeyspace := os.Getenv("DATABASE_KEYSPACE")
	cacheServer := os.Getenv("CACHE_SERVER")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	cassandra, err := database.CassandraConnect(databaseServer, databaseKeyspace)
	if err != nil {
		logger.Fatal(err)
	}
	defer cassandra.Close()

	redisCache := cache.NewRedisCache(cacheServer, 1, "")

	links := api.Links{
		Session:    cassandra,
		Cache:      redisCache,
		Expiration: time.Hour,
	}

	router := mux.NewRouter()
	router.HandleFunc("/{short_link}", links.Get).Methods(http.MethodGet)

	server.RunHTTP(host+":"+port, router)
}
