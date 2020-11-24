package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hjhussaini/url-shortener/models"
)

// swagger:route GET /{short_link}
func (links *Links) Get(writer http.ResponseWriter, request *http.Request) {
	var err error
	link := models.Link{}
	variables := mux.Vars(request)
	link.ShortURL = variables["short_link"]

	link.LongURL, err = links.Cache.Get(link.ShortURL)
	if err != nil {
		link.Select(links.Session)
	}

	if link.LongURL != "" {
		http.Redirect(writer, request, link.LongURL, 302)
		links.Cache.Set(link.ShortURL, link.LongURL, links.Expiration)

		return
	}

	http.Redirect(writer, request, "", http.StatusNotFound)
}
