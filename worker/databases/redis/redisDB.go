package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var rdb *redis.Client

func Connect(address string) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
}

func Set(k string, v string) error {
	return rdb.Set(ctx, k, v, 0).Err()
}

func HSet(k string, v ...interface{}) error {
	return rdb.HSet(ctx, k, v).Err()
}

func Get(k string) (string, error) {
	return rdb.Get(ctx, k).Result()
}

func Subscribe(channel string) *redis.PubSub {
	return rdb.Subscribe(ctx, channel)
}

func GetCtx() context.Context {
	return ctx
}
