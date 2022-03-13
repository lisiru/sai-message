package redisutil

import (
	"sai/cache"
	"sai/cache/redis"
)

var RedisClient cache.Factory

func NewRedisClient() cache.Factory {
	RedisClient, _ = redis.NewRedisFactoryOr(nil)
	return RedisClient
}
