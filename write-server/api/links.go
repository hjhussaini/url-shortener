package api

import (
	"net/http"

	"github.com/hjhussaini/url-shortener/cache"
	"github.com/hjhussaini/url-shortener/common"
	"github.com/hjhussaini/url-shortener/database"
	"github.com/hjhussaini/url-shortener/logger"
	"github.com/hjhussaini/url-shortener/models"
)

type Links struct {
	Session database.Session
	Cache   cache.Cache
}

func (links *Links) Caching() {
	for index := 0; index < 50; index++ {
		response, err := http.Get("http://keygen.server.com:5100/keys")
		if err != nil {
			logger.Error(err)

			return
		}
		defer response.Body.Close()

		if response.StatusCode == http.StatusNotFound {
			logger.Error(response.Status)

			return
		}

		data := &models.Data{}
		if err = common.ReadJSON(response.Body, data); err != nil {
			logger.Error(err)
			continue
		}

		if response.StatusCode != http.StatusOK {
			logger.Error(data.Message)
			continue
		}

		if err = links.Cache.Push(data.Key); err != nil {
			logger.Error(err)
		}
	}
}
