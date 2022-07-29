package main

import (
	"flag"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

var r = chi.NewRouter()

func init() {

	initLogger()
	container := InitializeServer()

	r.Route("/botConfig", func(r chi.Router) {
		r.Post("/save", container.BotConfigRedisService.SaveBotConfigRedis)
		r.Get("/get", container.BotConfigRedisService.GetBotConfigRedis)
	})
}

func initLogger() {
	zerolog.MessageFieldName = "message"
	zerolog.TimestampFieldName = "date"
	zerolog.LevelFieldName = "type"
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05 Z0700 UTC"
	log.Logger = log.With().Str("app", os.Getenv("APP_NAME")).Logger()
	debug := flag.Bool("debug", viper.GetBool("logger.debug"), "sets log level to debug")

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if e := log.Debug(); e.Enabled() {
		e.Msg("Debug mode enabled")
	}
}

func main() {

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Error().Str("method", "main").Msgf("%v", err)
	}
}
