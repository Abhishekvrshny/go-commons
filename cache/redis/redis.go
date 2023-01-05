package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Client holds redis client
type Client struct {
	client *redis.Client
}

// NewClient returns a new redis client
func NewClient(config Config) *Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.DB,
	})
	return &Client{client: client}
}

// Publish message to redis
func (c *Client) Publish(ctx context.Context, channel string, data []byte) error {
	err := c.client.Publish(ctx, channel, data).Err()
	if err != nil {
		return err
	}
	return nil
}

// Subscribe to a redis pubsub channel
func (c *Client) Subscribe(ctx context.Context, channel string) *redis.PubSub {
	return c.client.Subscribe(ctx, channel)
}

// UnSubscribe from a channel
func (c *Client) UnSubscribe(ctx context.Context, subs *redis.PubSub) error {
	err := subs.Close()
	if err != nil {
		return err
	}
	return nil
}
