package api

import (
	"net/http"

	"github.com/hjhussaini/url-shortener/common"
	"github.com/hjhussaini/url-shortener/models"
)

// swagger:route DELETE /
func (links *Links) DeleteURL(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")

	link := &models.Link{}
	if err := common.ReadJSON(request.Body, link); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		common.WriteJSON(models.Data{Message: err.Error()}, writer)

		return
	}

	if err := link.Delete(links.Session); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		common.WriteJSON(models.Data{Message: err.Error()}, writer)

		return
	}
}
