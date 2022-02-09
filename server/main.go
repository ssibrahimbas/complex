package main

import (
	"github.com/go-pg/pg/v10"
	redisDB "server/databases/redis"
	"server/handlers/testHandler"
	"server/handlers/valueHandler"
	"server/helpers/httpHelper"

	"github.com/gin-gonic/gin"
	"log"
	"server/config"
	"server/databases/postgre"
	"server/models/value"
)

func main() {
	config.LoadConfig()
	connectDatabases()
	initModels()
	r := gin.Default()
	initRoutes(r)
	log.Fatal(r.Run(config.Config.ServerHost + ":" + config.Config.ServerPort))
}

func connectDatabases() {
	postgre.Connect(&pg.Options{
		Addr:     config.Config.PgHost + ":" + config.Config.PgPort,
		Password: config.Config.PgPassword,
		User:     config.Config.PgUser,
		Database: config.Config.PgDatabase,
	})
	redisDB.Connect(config.Config.RedisHost + ":" + config.Config.RedisPort)
}

func initModels() {
	value.Init()
}

func initRoutes(r *gin.Engine) {
	r.Use(httpHelper.ErrorWrapper())
	r.GET("/", testHandler.Test)
	r.GET("/values/all", valueHandler.GetAll)
	r.GET("/values/current", valueHandler.GetCurrent)
	r.POST("/values", valueHandler.Create)
}
