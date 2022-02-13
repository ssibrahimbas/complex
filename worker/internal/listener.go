package internal

import (
	"complex/lib/math"
	redisDB "complex/tools/redis"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type Listener struct {
	redisClient *redisDB.RedisClient
	math        *math.Math
}

func New(redisClient *redisDB.RedisClient, math *math.Math) *Listener {
	return &Listener{
		math:        math,
		redisClient: redisClient,
	}
}

func (l *Listener) SubscribeAndListen() {
	topic := l.Subscribe()
	channel := l.GetChannel(topic)
	l.Listen(channel)
}

func (l *Listener) Subscribe() *redis.PubSub {
	return l.redisClient.Subscribe("insert")
}

func (l *Listener) GetChannel(p *redis.PubSub) <-chan *redis.Message {
	return p.Channel()
}

func (l *Listener) Listen(channel <-chan *redis.Message) {
	for msg := range channel {
		err := l.OnMessageReceive(msg.Payload)
		if err != nil {
			panic(err)
		}
	}
}

func (l *Listener) OnMessageReceive(msg string) error {
	ind, err := strconv.Atoi(msg)
	if err != nil {
		return err
	}
	l.redisClient.HSet("values", msg, strconv.Itoa(l.math.Fib(ind)))
	return nil
}
