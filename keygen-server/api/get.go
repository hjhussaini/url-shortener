package api

import (
	"net/http"

	"github.com/hjhussaini/url-shortener/common"
	"github.com/hjhussaini/url-shortener/logger"
	"github.com/hjhussaini/url-shortener/models"
)

// swagger:route GET /keys
func (api *API) GetKey(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		return
	}

	api.mutex.Lock()
	defer api.mutex.Unlock()

	key, err := api.Cache.Pop()
	if err != nil {
		api.Caching()
		key, err = api.Cache.Pop()
		if err != nil {
			logger.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
			common.WriteJSON(models.Data{Message: err.Error()}, writer)

			return
		}
	}
	common.WriteJSON(models.Data{Key: key}, writer)
}
