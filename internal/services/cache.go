package services

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type CacheService struct {
	client *redis.Client
}

func NewCacheService(addr string) *CacheService {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &CacheService{client: client}
}

func (c *CacheService) Set(key string, value interface{}) error {
	return c.client.Set(ctx, key, value, 0).Err()
}

func (c *CacheService) Get(key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *CacheService) Invalidate(key string) error {
	return c.client.Del(ctx, key).Err()
}
