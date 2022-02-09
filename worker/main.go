package main

import (
	"complex/config"
	redisDB "complex/databases/redis"
	"complex/lib/math"
	"fmt"
	"os"
	"strconv"
)

func main() {
	err := os.Setenv("REDIS_HOST", "localhost")
	if err != nil {
		return
	}
	err2 := os.Setenv("REDIS_PORT", "6379")
	if err2 != nil {
		return
	}
	config.LoadConfig()
	redisAddress := config.Config.RedisHost + ":" + config.Config.RedisPort
	redisDB.Connect(redisAddress)
	listenAndSetFib()
}

func listenAndSetFib() {
	sub := redisDB.Subscribe("insert")
	msg, err := sub.ReceiveMessage(redisDB.GetCtx())
	if err != nil {
		fmt.Println("err : " + err.Error())
		panic(err)
	}
	num, err := strconv.Atoi(msg.Payload)
	if err != nil {
		panic(err)
	}
	redisDB.HSet("value", strconv.Itoa(math.Fib(num)))
}
