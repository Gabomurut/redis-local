//go:build InitializeServer
// +build InitializeServer

package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
	"github.com/redisBotConfig/client"
	"github.com/redisBotConfig/config"
	"github.com/redisBotConfig/pkg/service"
)

func InitializeServer() config.ContainerService {
	wire.Build(
		resty.New,
		service.NewBotConfigServiceImpl,
		config.NewContainerService,
		client.RedisClientImpl,
	)

	return config.ContainerService{}
}
