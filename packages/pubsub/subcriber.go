package pubsub

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type Subscription[Event any] interface {
	Channel() <-chan Event
	Close() error
}

type Subscriber[Event any] interface {
	Subscribe(ctx context.Context, topic string) (Subscription[Event], error)
}

type RedisSubscriber[Event any] struct {
	redis *redis.Client
}

type RedisSubscription[Event any] struct {
	sub *redis.PubSub
	channel <-chan Event
}

func NewRedisSubscriber[Event any](redis *redis.Client) Subscriber[Event] {
	return &RedisSubscriber[Event]{
		redis: redis,
	}
}

func (s *RedisSubscriber[Event]) Subscribe(ctx context.Context, topic string) (Subscription[Event], error) {
	sub := s.redis.Subscribe(ctx, topic)
	ch := make(chan Event)


	go func() {
		defer close(ch)
		for {
			select {
			case <- ctx.Done():
				return
			case msg := <- sub.Channel():
				var event Event
				err := json.Unmarshal([]byte(msg.Payload), &event)
				if err != nil {
					return
				}
				ch <- event
			}
		}
	}()

	return &RedisSubscription[Event]{sub: sub, channel: ch}, nil
}

func (s *RedisSubscription[Event]) Channel() <-chan Event {
	return s.channel
}

func (s *RedisSubscription[Event]) Close() error {
	return s.sub.Close()
}
