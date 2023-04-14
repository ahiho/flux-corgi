package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"

	log "fluxcorgi/pkg/logger"

	"fluxcorgi/internal/server"
)

func main() {
	var logger zerolog.Logger

	if os.Getenv("ENV") == "DEVELOPMENT" {
		err := godotenv.Load()
		if err != nil {
			logger.Fatal().Msgf("Error loading ENV: %v", err)
		}
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05.000000"}).With().Timestamp().Logger()
	} else {
		zerolog.TimeFieldFormat = "2006-01-02T15:04:05.999999Z07:00"
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	}

	log.SetGlobalLogger(logger)

	err := server.RunServer()
	if err != nil {
		logger.Fatal().Msgf("error: %v\n", err.Error())
	}
}
