package cache

import (
	"context"
	"log"
	"sync"

	"github.com/nikhilryan/go-featuristic/config"
	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

func GetRedisClient() *redis.Client {
	redisOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr: config.GetRedisAddr(),
		})

		_, err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalf("failed to connect to Redis: %v", err)
		}
	})
	return redisClient
}
