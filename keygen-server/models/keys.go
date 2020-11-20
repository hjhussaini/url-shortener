package models

import (
	"github.com/hjhussaini/url-shortener/database"
)

// sqagger:model
type Keys struct {
	// The unique key
	Key string
}

func (keys *Keys) Table() string {
	return "keys"
}

func (keys *Keys) Fields() string {
	return "key"
}

func (keys *Keys) Count(session database.Session) int64 {
	count, _ := session.Count(keys.Table())

	return count
}

func (keys *Keys) Insert(session database.Session) error {
	return session.Insert(0, keys.Table(), keys.Fields(), keys.Key)
}
