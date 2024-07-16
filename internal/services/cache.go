package services

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type CacheService struct {
	client *redis.Client
}

func NewAppCacheService(client *redis.Client) *CacheService {
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
