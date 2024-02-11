package main

import (
	"context"
	"log"

	"github.com/slvrblue/cbt-bot/internal/app"
	"github.com/slvrblue/cbt-bot/internal/config"
)

func main() {
	cfg, err := config.New[config.Config]()
	if err != nil {
		log.Fatalf("error parsing config: %s\n", err)
	}

	a, err := app.New(cfg)
	if err != nil {
		log.Fatalf("error creating app: %s\n", err)
	}

	if err := a.Run(context.Background()); err != nil {
		log.Fatalf("error running app: %s\n", err)
	}
}
