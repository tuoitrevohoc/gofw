package pubsub

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Publisher[Event any] interface {
	Publish(ctx context.Context, topic string, message Event) error
}

type RedisPublisher[Event any] struct {
	redis *redis.Client
}

func NewRedisPublisher[Event any](redis *redis.Client) Publisher[Event] {
	return &RedisPublisher[Event]{
		redis: redis,
	}
}

func (p *RedisPublisher[Event]) Publish(ctx context.Context, topic string, message Event) error {
	return p.redis.Publish(ctx, topic, message).Err()
}
