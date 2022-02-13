package main

import (
	"server/config"
	"server/http"
	values "server/internal"
	"server/tools/databases"
	"server/tools/redis"
)

func main() {
	conf := config.LoadConfig()
	dbs := databases.NewClient(conf)
	pg := dbs.ConnectPostgres()
	rdb := dbs.ConnectRedis()
	redisClient := redis.NewClient(rdb)

	repo := values.NewRepository(pg)
	service := values.NewService(repo, redisClient)
	handler := values.NewHandler(service)

	err := http.NewServer(handler.Init(), conf.ServerHost, conf.ServerPort).Run()
	if err != nil {
		panic("Couldn't start the HTTP server.")
	}
}
