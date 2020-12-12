package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hjhussaini/url-shortener/cache"
	"github.com/hjhussaini/url-shortener/database"
	"github.com/hjhussaini/url-shortener/link-server/api"
	"github.com/hjhussaini/url-shortener/server"
)

func main() {
	databaseServer := os.Getenv("DATABASE_SERVER")
	databaseKeyspace := os.Getenv("DATABASE_KEYSPACE")
	cacheServer := os.Getenv("CACHE_SERVER")
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	cassandra, err := database.CassandraConnect(databaseServer, databaseKeyspace)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer cassandra.Close()

	redisCache := cache.NewRedisCache(cacheServer, 1, "keys_cache")

	links := api.Links{
		Session: cassandra,
		Cache:   redisCache,
	}

	router := mux.NewRouter()
	router.HandleFunc("/{short_link}", links.RedirectURL).Methods(http.MethodGet)
	router.HandleFunc("/links", links.CreateURL).Methods(http.MethodPost)
	router.HandleFunc("/links/{short_link}", links.DeleteURL).Methods(http.MethodDelete)

	log.Printf("Serving HTTP server on 0.0.0.0:%d\n", port)
	if err := server.RunHTTP(port, router); err != nil {
		log.Fatal(err)
	}
}
