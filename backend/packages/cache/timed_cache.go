package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type TimedCache interface {
	Get(ctx context.Context, key string, value interface{}) error
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
}

type timedCache struct {
	// redis client
	redis *redis.Client
}

func NewTimedCache(redis *redis.Client) TimedCache {
	return &timedCache{
		redis: redis,
	}
}

func (c *timedCache) Get(ctx context.Context, key string, value interface{}) error {
	result, err := c.redis.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(result), value)
	if err != nil {
		return err
	}

	return nil
}

func (c *timedCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.redis.Set(ctx, key, json, ttl).Err()
}

func (c *timedCache) Delete(ctx context.Context, key string) error {
	return c.redis.Del(context.Background(), key).Err()
}
