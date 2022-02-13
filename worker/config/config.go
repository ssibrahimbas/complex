package config

import "os"

type Configuration struct {
	RedisHost string
	RedisPort string
}

func LoadConfig() *Configuration {
	return &Configuration{
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
	}
}
