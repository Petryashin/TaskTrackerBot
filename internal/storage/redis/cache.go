package rediscache

import (
	redis "github.com/go-redis/redis"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) RedisCache {
	return RedisCache{client: client}
}

func (rc RedisCache) Set(key string, json string) error {
	return rc.client.Set(key, json, 0).Err()
}

func (rc RedisCache) Get(key string) (string, error) {
	return rc.client.Get(key).Result()
}
