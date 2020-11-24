package api

import (
	"fmt"
	"net/http"

	"github.com/hjhussaini/url-shortener/common"
	"github.com/hjhussaini/url-shortener/models"
)

// swagger:route POST /
func (links *Links) CreateURL(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")

	link := &models.Link{}
	err := common.ReadJSON(request.Body, link)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	link.ShortURL, err = links.Cache.Pop()
	if err != nil {
		links.Caching()
		link.ShortURL, err = links.Cache.Pop()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			common.WriteJSON(models.Data{Message: err.Error()}, writer)

			return
		}
	}
	if err := link.Insert(links.Session); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(writer, err.Error())
		return
	}
}
