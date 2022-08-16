package service

import (
	"encoding/json"
	client2 "github.com/redisBotConfig/pkg/client"
	"github.com/redisBotConfig/pkg/dto"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

const (
	keyBotConfig = "ms-portal-bot-botConfig-sandbox-did-"
	keyBotToken  = "ms-portal-bot-"
	keyBotToken2 = "-sandbox-botToken"
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
	//save value on redis
	redisClient := client2.RedisClient
	var redisBot dto.RedisBot
	err := json.NewDecoder(r.Body).Decode(&redisBot)

	if err != nil {
		log.Error().Msgf("error decoding json {}", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	did := strings.Split(redisBot.Token, ":")
	log.Info().Msgf("DID {}", did[0])

	response, err2 := json.Marshal(redisBot.BotConfig)
	log.Info().Msg("BotConfig Marshalled")

	if err2 != nil {
		log.Error().Msgf("error marshalling json to BotConfig dto {}", err2)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	botKey := keyBotConfig + did[0]
	err = redisClient().Set(botKey, response, 0).Err()

	if err != nil {
		log.Error().Msgf("Error setting BotConfig value on Redis {}", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tokenKey := keyBotToken + did[0] + keyBotToken2

	err = redisClient().Set(tokenKey, redisBot.Token, 0).Err()
	if err != nil {
		log.Error().Msgf("Error setting token value on Redis {}", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (b *BotConfigRedisServiceImpl) GetBotConfigRedis(w http.ResponseWriter, r *http.Request) {
	// get value from Redis
	redisClient := client2.RedisClient

	response, err := redisClient().Get("ms-portal-bot-botConfig-sandbox-did-123456789").Result()
	if err != nil {
		log.Error().Msgf("Error getting value on Redis {}", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
