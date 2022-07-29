package service

import (
	"fmt"
	client2 "github.com/redisBotConfig/client"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

// BotConfigRedisServiceImpl - ServiceImpl struct definition
type BotConfigRedisServiceImpl struct {
	log zerolog.Logger
}

// NewBotConfigServiceImpl - constructor
func NewBotConfigServiceImpl() BotConfigRedisService {
	return &BotConfigRedisServiceImpl{}
}

func (b *BotConfigRedisServiceImpl) SaveBotConfigRedis(w http.ResponseWriter, r *http.Request) {

	redisClient := client2.RedisClient
	err := redisClient().Set("PRUBAFINAL", "PRUEBA", 0).Err()

	if err != nil {
		log.Fatal()
	}
	fmt.Println("TESTED")
}

func (b *BotConfigRedisServiceImpl) GetBotConfigRedis(w http.ResponseWriter, r *http.Request) {
	// get value
	redisClient := client2.RedisClient
	username, err := redisClient().Get("TESTEO").Result()
	if err != nil {
		log.Fatal()
	}
	fmt.Println(username)
}
