package pubsub

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type Subscription interface {
	Channel() <-chan string
	Close() error
}

type Subscriber interface {
	Subscribe(ctx context.Context, topic string) (Subscription, error)
}

type RedisSubscriber struct {
	redis *redis.Client
}

type RedisSubscription struct {
	sub *redis.PubSub
	channel <-chan string
}

func NewRedisSubscriber(redis *redis.Client) Subscriber {
	return &RedisSubscriber{
		redis: redis,
	}
}

func (s *RedisSubscriber) Subscribe(ctx context.Context, topic string) (Subscription, error) {
	sub := s.redis.Subscribe(ctx, topic)
	ch := make(chan string)


	go func() {
		defer close(ch)
		for {
			select {
			case <- ctx.Done():
				return
			case msg := <- sub.Channel():
				ch <- msg.Payload
			}
		}
	}()

	return &RedisSubscription{sub: sub, channel: ch}, nil
}

func SubscribeToChannel[Event any](ctx context.Context, subscriber Subscriber, topic string) (<-chan Event, error) {
	sub, err := subscriber.Subscribe(ctx, topic)
	if err != nil {
		return nil, err
	}
	ch := make(chan Event)

	go func() {	
		defer close(ch)
		for {
			select {
			case <- ctx.Done():
				sub.Close()
				return
			case msg := <- sub.Channel():
				var event Event
				err := json.Unmarshal([]byte(msg), &event)
				if err != nil {
					return
				}
				ch <- event
			}
		}
	}()

	return ch, nil
}

func (s *RedisSubscription) Close() error {
	return s.sub.Close()
}

func (s *RedisSubscription) Channel() <-chan string {
	return s.channel
}

