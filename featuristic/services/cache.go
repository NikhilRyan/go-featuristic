package services

import (
	"context"
	"sync"
)

type CacheService struct {
	client RedisClient
	mu     sync.Mutex
}

func NewAppCacheService(client RedisClient) *CacheService {
	return &CacheService{
		client: client,
	}
}

func (c *CacheService) Set(key string, value string) error {
	return c.client.Set(context.Background(), key, value, 0)
}

func (c *CacheService) Get(key string) (string, error) {
	return c.client.Get(context.Background(), key)
}

func (c *CacheService) Delete(key string) error {
	return c.client.Del(context.Background(), key)
}

func (c *CacheService) DeleteNamespace(namespace string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	keys, err := c.client.Keys(context.Background(), namespace+"_*")
	if err != nil {
		return err
	}

	if len(keys) > 0 {
		return c.client.Del(context.Background(), keys...)
	}

	return nil
}
