package client

import (
	"bitbucket.org/chattigodev/chattigo-golang-library/pkg/utils"
	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

func RedisClient() *redis.Client {
	sublogger := log.With().Str(utils.Method, "RedisClient").Logger()
	sublogger.Info().Msg(utils.InitStr)
	database, err := strconv.Atoi(os.Getenv("DATABASE"))
	if err != nil {
		sublogger.Err(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("HOST") + os.Getenv("PORT"),
		Password: os.Getenv("PASSWORD"),
		DB:       database,
	})
	return client
}
