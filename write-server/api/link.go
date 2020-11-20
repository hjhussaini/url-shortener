package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hjhussaini/url-shortener/database"
)

type Link struct {
	Session database.Session
}

func (link *Link) Register(router *mux.Router) {
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", nil)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/", nil)
}

func NewLinkAPIs(session database.Session) *Link {
	return &Link{Session: session}
}
