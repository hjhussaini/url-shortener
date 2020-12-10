package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	Name    string
	Client  *redis.Client
}

func (cache *RedisCache) Set(
	key string,
	value interface{},
	expiration time.Duration,
) error {
	return cache.Client.Set(context.Background(), key, value, expiration).Err()
}

func (cache *RedisCache) Get(key string) (string, error) {
	return cache.Client.Get(context.Background(), key).Result()
}

func (cache *RedisCache) Push(values ...interface{}) error {
	return cache.Client.LPush(context.Background(), cache.Name, values...).Err()
}

func (cache *RedisCache) Pop() (string, error) {
	return cache.Client.RPop(context.Background(), cache.Name).Result()
}

func NewRedisCache(address string, database int, name string) *RedisCache {
	cache := &RedisCache{Name: name}
	cache.Client = redis.NewClient(&redis.Options{
		Addr: address,
		DB:   database,
	})

	return cache
}
