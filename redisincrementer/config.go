package redisincrementer

import "github.com/crosscode-nl/incwebdemo/env"

type Config struct {
	RedisAddr string
}

func NewConfig() *Config {
	return &Config{
		RedisAddr: env.Get("REDISADDR","localhost:6379"),
	}
}

