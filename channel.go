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

func (c *RedisChannelManger) Subscribe(ctx context.Context, channel string) *redis.PubSub {
	return c.rd.Subscribe(ctx, channel)
}

func (c *RedisChannelManger) Publish(ctx context.Context, channel string, message any) error {
	bs, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("marshal message: %w", err)
	}

	err = c.rd.Publish(ctx, channel, bs).Err()
	if err != nil {
		return fmt.Errorf("publish message: %w", err)
	}
	return nil
}
