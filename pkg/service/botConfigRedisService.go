package service

import "net/http"

type BotConfigRedisService interface {
	SaveBotConfigRedis(w http.ResponseWriter, r *http.Request)
	GetBotConfigRedis(w http.ResponseWriter, r *http.Request)
}
