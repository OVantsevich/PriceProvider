// Package repository repos
package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"Price-Provider/internal/model"

	"github.com/go-redis/redis/v8"
)

// Redis entity
type Redis struct {
	Client     *redis.Client
	StreamName string
}

// NewRedis constructor
func NewRedis(client *redis.Client, streamName string) *Redis {
	return &Redis{Client: client, StreamName: streamName}
}

// PublishPrices publish prices to the stream
func (c *Redis) PublishPrices(ctx context.Context, prices []*model.Price) error {
	mp, _ := json.Marshal(prices)
	_, err := c.Client.XAdd(ctx, &redis.XAddArgs{
		Stream: c.StreamName,
		Values: map[string]interface{}{
			"prices": mp,
		},
	}).Result()
	if err != nil {
		return fmt.Errorf("redis - PublishPrices - XAdd: %w", err)
	}
	return nil
}
