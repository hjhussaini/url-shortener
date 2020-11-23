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

func (usedKeys *UsedKeys) Fields() string {
	return "key"
}

func (usedKeys *UsedKeys) Count() int64 {
	count, _ := usedKeys.Session.Count(usedKeys.Table())

	return count
}

func (usedKeys *UsedKeys) Insert() error {
	return usedKeys.Session.Insert(
		0,
		usedKeys.Table(),
		usedKeys.Fields(),
		usedKeys.Key,
	)
}
