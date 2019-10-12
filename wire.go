// +build wireinject

package main

import (
	"github.com/crosscode-nl/incwebdemo/application"
	"github.com/crosscode-nl/incwebdemo/incrementer"
	"github.com/crosscode-nl/incwebdemo/memoryincrementer"
	"github.com/crosscode-nl/incwebdemo/redisincrementer"
	"github.com/google/wire"
)

func InitializeMemoryApp() *application.App {
	wire.Build(memoryincrementer.New, wire.Bind(new(incrementer.Incrementer),new(*memoryincrementer.MemIncrementer)), application.New)
	return &application.App{}
}


func InitializeRedisApp() *application.App {
	wire.Build(redisincrementer.NewRedisIncrementer, redisincrementer.NewConfig, wire.Bind(new(incrementer.Incrementer),new(*redisincrementer.RedisIncrementer)), application.New)
	return &application.App{}
}
