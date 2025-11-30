package main

import (
	"company-api/app/api/app"
	"company-api/app/api/handler"
	"company-api/app/api/route"
	"company-api/foundation/logger"
	"context"
	"fmt"
	"log"
	"net/http"
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

	h := handler.New(logger)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), route.Routes(h)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Println("Application started successfully")
}
