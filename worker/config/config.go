package config

import "os"

type Configuration struct {
	RedisHost string
	RedisPort string
}

var Config Configuration

func LoadConfig() {
	Config = Configuration{
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
	}
}
