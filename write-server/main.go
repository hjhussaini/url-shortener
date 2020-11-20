package main

import (
	"fmt"
	"os"

	"github.com/hjhussaini/url-shortener/database"
	"github.com/hjhussaini/url-shortener/server"
)

func main() {
	databaseServer := os.Getenv("DATABASE_SERVER")
	databaseKeyspace := os.Getenv("DATABASE_KEYSPACE")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	cassandra, err := database.CassandraConnect(databaseServer, databaseKeyspace)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer cassandra.Close()

	server.RunHTTP(host+":"+port, nil)
}
