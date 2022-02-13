package main

import (
	"complex/config"
	"complex/internal"
	"complex/lib/math"
	"complex/tools/databases"
	redisDB "complex/tools/redis"
)

func main() {
	conf := config.LoadConfig()
	dbs := databases.NewClient(conf)
	rdb := dbs.ConnectRedis()
	redis := redisDB.NewClient(rdb)
	_math := math.New()
	listener := internal.New(redis, _math)
	listener.SubscribeAndListen()
}
