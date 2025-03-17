package pubsub

import (
	"context"

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
	return p.redis.Publish(ctx, topic, message).Err()
}
