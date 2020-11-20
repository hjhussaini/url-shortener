package main

import (
	"os"

	"github.com/hjhussaini/url-shortener/server"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	server.RunHTTP(host+":"+port, nil)
}
