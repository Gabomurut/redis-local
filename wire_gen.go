// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/redisBotConfig/config"
	"github.com/redisBotConfig/service"
)

// Injectors from wire.go:

func InitializeServer() config.ContainerService {
	client := resty.New()
	botConfigRedisService := service.NewBotConfigServiceImpl()
	containerService := config.NewContainerService(client, botConfigRedisService)
	return containerService
}