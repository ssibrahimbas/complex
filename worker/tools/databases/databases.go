package databases

import (
	"complex/config"
	"github.com/go-redis/redis/v8"
)

type Databases struct {
	config *config.Configuration
}

func NewClient(config *config.Configuration) *Databases {
	return &Databases{
		config: config,
	}
}

func (d Databases) ConnectRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     d.config.RedisHost + ":" + d.config.RedisPort,
		Password: "",
		DB:       0,
	})
}
