package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hjhussaini/url-shortener/cache"
	"github.com/hjhussaini/url-shortener/database"
	"github.com/hjhussaini/url-shortener/server"
	"github.com/hjhussaini/url-shortener/write-server/api"
)

func main() {
	databaseServer := os.Getenv("DATABASE_SERVER")
	databaseKeyspace := os.Getenv("DATABASE_KEYSPACE")
	cacheServer := os.Getenv("CACHE_SERVER")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	cassandra, err := database.CassandraConnect(databaseServer, databaseKeyspace)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer cassandra.Close()

	redisCache := cache.NewRedisCache(cacheServer, 2, "keys_cache")
	links := &api.Links{
		Session: cassandra,
		Cache:   redisCache,
	}
	links.Caching()

	router := mux.NewRouter()
	router.HandleFunc("/", links.CreateURL).Methods(http.MethodPost)

	server.RunHTTP(host+":"+port, router)
}
