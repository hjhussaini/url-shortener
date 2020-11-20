package models

import "github.com/hjhussaini/url-shortener/database"

// swagger:model
type UsedKeys struct {
	// The used unique key
	Key string
}

func (usedKeys *UsedKeys) Table() string {
	return "used_keys"
}

func (usedKeys *UsedKeys) Count(session database.Session) int64 {
	count, _ := session.Count(usedKeys.Table())

	return count
}
