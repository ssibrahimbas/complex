package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisClient struct {
	client *redis.Client
}

func NewClient(client *redis.Client) *RedisClient {
	return &RedisClient{
		client: client,
	}
}

func (r *RedisClient) Set(k string, v string) error {
	return r.client.Set(ctx, k, v, 0).Err()
}

func (r *RedisClient) HSet(k string, v ...interface{}) error {
	return r.client.HSet(ctx, k, v).Err()
}

func (r *RedisClient) Get(k string) (string, error) {
	return r.client.Get(ctx, k).Result()
}

func (r *RedisClient) HGetAll(k string) (map[string]string, error) {
	return r.client.HGetAll(ctx, k).Result()
}

func (r *RedisClient) Subscribe(channel string) *redis.PubSub {
	return r.client.Subscribe(ctx, channel)
}

func (r *RedisClient) Publish(ch string, val string) {
	r.client.Publish(ctx, ch, val)
}
