package main

import (
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
	"restaurant/internal/statistics_statistics/app"
	"restaurant/internal/statistics_statistics/config"
)

func main() {
	cfg := config.Config{}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to retrieve env variables, %v", err)
	}

	if err := app.Run(cfg); err != nil {
		log.Fatal("error running server ", err)
	}
}
