package cache

import (
	"context"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	Client  *redis.Client
	Context context.Context
}

func (cache *RedisCache) Push(key string, value string) error {
	return cache.Client.LPush(cache.Context, key, value).Err()
}

func (cache *RedisCache) Pop(key string) (string, error) {
	return cache.Client.LPop(cache.Context, key).Result()
}

func NewRedisCache(address string, database int) *RedisCache {
	cache := &RedisCache{}
	cache.Client = redis.NewClient(&redis.Options{
		Addr: address,
		DB:   database,
	})
	cache.Context = context.Background()

	return cache
}
