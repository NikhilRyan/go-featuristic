package services

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// RedisUniversalClientAdapter adapts redis.UniversalClient to RedisClient interface.
type RedisUniversalClientAdapter struct {
	client redis.UniversalClient
}

func NewRedisUniversalClientAdapter(client redis.UniversalClient) *RedisUniversalClientAdapter {
	return &RedisUniversalClientAdapter{client: client}
}

func (r *RedisUniversalClientAdapter) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisUniversalClientAdapter) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisUniversalClientAdapter) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

func (r *RedisUniversalClientAdapter) Keys(ctx context.Context, pattern string) ([]string, error) {
	return r.client.Keys(ctx, pattern).Result()
}

// RedisClientAdapter adapts redis.Client to RedisClient interface.
type RedisClientAdapter struct {
	client *redis.Client
}

func NewRedisClientAdapter(client *redis.Client) *RedisClientAdapter {
	return &RedisClientAdapter{client: client}
}

func (r *RedisClientAdapter) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisClientAdapter) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisClientAdapter) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

func (r *RedisClientAdapter) Keys(ctx context.Context, pattern string) ([]string, error) {
	return r.client.Keys(ctx, pattern).Result()
}
