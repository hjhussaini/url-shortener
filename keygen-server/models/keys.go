package models

import (
	"github.com/hjhussaini/url-shortener/database"
)

// sqagger:model
type Keys struct {
	Session database.Session
	// The unique key
	Key string
}

func (keys *Keys) Table() string {
	return "keys"
}

func (keys *Keys) Fields() string {
	return "key"
}

func (keys *Keys) Count() int64 {
	count, _ := keys.Session.Count(keys.Table())

	return count
}

func (keys *Keys) Select(count int) []string {
	var value string
	var values []string

	result := keys.Session.Select(keys.Table(), keys.Fields(), "", count)
	for result.Scan(&value) {
		values = append(values, value)
	}

	return values
}

func (keys *Keys) Insert() error {
	return keys.Session.Insert(0, keys.Table(), keys.Fields(), keys.Key)
}
