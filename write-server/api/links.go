package api

import (
	"github.com/hjhussaini/url-shortener/cache"
	"github.com/hjhussaini/url-shortener/database"
)

type Links struct {
	Session database.Session
	Cache   cache.Cache
}
