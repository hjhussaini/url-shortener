package models

import "github.com/hjhussaini/url-shortener/database"

// swagger:model
type UsedKeys struct {
	Session database.Session
	// The used unique key
	Key string
}

func (usedKeys *UsedKeys) Table() string {
	return "used_keys"
}

func (usedKeys *UsedKeys) Count() int64 {
	count, _ := usedKeys.Session.Count(usedKeys.Table())

	return count
}
