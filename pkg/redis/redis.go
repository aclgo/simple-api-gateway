package redis

import "github.com/go-redis/redis/v8"

func NewRedisClient() *redis.Client {
	return &redis.Client{}
}
