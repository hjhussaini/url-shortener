package cache

import (
	"context"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	Name    string
	Client  *redis.Client
	Context context.Context
}

func (cache *RedisCache) Push(values ...interface{}) error {
	return cache.Client.LPush(cache.Context, cache.Name, values...).Err()
}

func (cache *RedisCache) Pop() (string, error) {
	return cache.Client.RPop(cache.Context, cache.Name).Result()
}

func NewRedisCache(address string, database int, name string) *RedisCache {
	cache := &RedisCache{Name: name}
	cache.Client = redis.NewClient(&redis.Options{
		Addr: address,
		DB:   database,
	})
	cache.Context = context.Background()

	return cache
}
