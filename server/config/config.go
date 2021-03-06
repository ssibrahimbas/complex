package config

import (
	"fmt"
	"os"
)

type Configuration struct {
	RedisHost  string
	RedisPort  string
	PgUser     string
	PgHost     string
	PgDatabase string
	PgPassword string
	PgPort     string
	ServerHost string
	ServerPort string
}

func LoadConfig() *Configuration {
	fmt.Println(os.Getenv("PG_HOST"))
	return &Configuration{
		RedisHost:  os.Getenv("REDIS_HOST"),
		RedisPort:  os.Getenv("REDIS_PORT"),
		PgUser:     os.Getenv("PG_USER"),
		PgHost:     os.Getenv("PG_HOST"),
		PgDatabase: os.Getenv("PG_DATABASE"),
		PgPassword: os.Getenv("PG_PASSWORD"),
		PgPort:     os.Getenv("PG_PORT"),
		ServerHost: os.Getenv("SERVER_HOST"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
