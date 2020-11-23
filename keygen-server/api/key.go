package api

import (
	"sync"

	"github.com/hjhussaini/url-shortener/cache"
	"github.com/hjhussaini/url-shortener/database"
)

type Keys struct {
	mutex   sync.Mutex
	Session database.Session
	Cache   cache.Cache
}
