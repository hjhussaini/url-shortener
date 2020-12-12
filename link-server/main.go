package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hjhussaini/url-shortener/link-server/api"
	"github.com/hjhussaini/url-shortener/server"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	links := api.Links{}

	router := mux.NewRouter()
	router.HandleFunc("/{short_link}", links.RedirectURL).Methods(http.MethodGet)
	router.HandleFunc("/links", links.CreateURL).Methods(http.MethodPost)
	router.HandleFunc("/links/{short_link}", links.DeleteURL).Methods(http.MethodDelete)

	log.Printf("Serving HTTP server on 0.0.0.0:%d\n", port)
	if err := server.RunHTTP(port, router); err != nil {
		log.Fatal(err)
	}
}
