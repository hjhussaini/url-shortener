package cache

import "time"

type Cache interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Push(value ...interface{}) error
	Pop() (string, error)
}
