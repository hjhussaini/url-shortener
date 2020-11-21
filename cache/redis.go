package cache

import "github.com/go-redis/redis"

type RedisCache struct {
	Client *redis.Client
}

func NewRedisCache(address string, database int) *RedisCache {
	cache := &RedisCache{}
	cache.Client = redis.NewClient(&redis.Options{
		Addr: address,
		DB:   database,
	})

	return cache
}
