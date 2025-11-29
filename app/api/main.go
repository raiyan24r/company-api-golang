package main

import (
	"company-api/app/api/app"
	"company-api/foundation/logger"
	"context"
	"log"
)

func main() {
	log.Println("Starting the application...")

	logger, err := logger.New()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	cfg, err := app.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	_ = app.New(context.Background(), *cfg, logger)
}
