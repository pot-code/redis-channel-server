package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisChannelManger struct {
	rd *redis.Client
}

func NewRedisChannelManger(rd *redis.Client) *RedisChannelManger {
	return &RedisChannelManger{
		rd: rd,
	}
}

func (c *RedisChannelManger) Subscribe(ctx context.Context, channel string) <-chan []byte {
	sub := c.rd.Subscribe(ctx, channel)
	ch := make(chan []byte)
	go func() {
		for msg := range sub.Channel() {
			bs, err := json.Marshal(&Message{
				Channel: channel,
				Message: msg.Payload,
			})
			if err != nil {
				panic(fmt.Errorf("marshal message: %w", err))
			}
			ch <- bs
		}
		close(ch)
	}()
	return ch
}

func (c *RedisChannelManger) Publish(ctx context.Context, message *Message) error {
	bs, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("marshal message: %w", err)
	}

	err = c.rd.Publish(ctx, message.Channel, bs).Err()
	if err != nil {
		return fmt.Errorf("publish message: %w", err)
	}
	return nil
}
