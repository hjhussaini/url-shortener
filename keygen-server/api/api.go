package api

import (
	"sync"

	"github.com/hjhussaini/url-shortener/cache"
	"github.com/hjhussaini/url-shortener/database"
	"github.com/hjhussaini/url-shortener/keygen-server/models"
	"github.com/hjhussaini/url-shortener/logger"
)

type API struct {
	mutex   sync.Mutex
	Session database.Session
	Cache   cache.Cache
}

func (api *API) Caching() error {
	keys := models.Keys{Session: api.Session}
	usedKeys := models.UsedKeys{Session: api.Session}
	values := keys.Select(50)
	for _, value := range values {
		usedKeys.Key = value
		if err := usedKeys.Insert(); err != nil {
			logger.Error(err)
			continue
		}
		keys.Key = usedKeys.Key
		if err := keys.Delete(); err != nil {
			logger.Error(err)
		}
	}

	return api.Cache.Push(values)
}
