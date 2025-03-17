package pubsub

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type Publisher interface {
	Publish(ctx context.Context, topic string, event interface{}) error
}

type RedisPublisher struct {
	redis *redis.Client
}

func NewRedisPublisher(redis *redis.Client) Publisher {
	return &RedisPublisher{
		redis: redis,
	}
}

func (p *RedisPublisher) Publish(ctx context.Context, topic string, message interface{}) error {
	// encode message to json
	json, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return p.redis.Publish(ctx, topic, json).Err()
}
