package config

import (
	"github.com/go-resty/resty/v2"
	"github.com/redisBotConfig/service"
)

type ContainerService struct {
	BotConfigRedisService service.BotConfigRedisService
	restyClient           *resty.Client
}

func NewContainerService(resty *resty.Client, botConfigRedisService service.BotConfigRedisService) ContainerService {

	resty.SetDoNotParseResponse(true)
	return ContainerService{
		BotConfigRedisService: botConfigRedisService,
	}
}
