package redis

import (
	"context"
	"log"

	"github.com/aclgo/simple-api-gateway/config"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("redis.Ping: %v", err)
	}

	return client
}
