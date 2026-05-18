package main

import (
	"company-api/app/api/app"
	"company-api/app/api/handler"
	"company-api/business/database"
	mysqldb "company-api/foundation/database"
	"company-api/foundation/logger"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Initializing application...")

	logger, err := logger.New()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	cfg, err := app.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	mysqlDb, _ := mysqldb.Open(cfg.DB)
	dbRepo := database.New(mysqlDb)

	a := app.New(context.Background(), *cfg, logger, dbRepo)

	h := handler.New(a)

	log.Println("Application started on port"+fmt.Sprintf(":%d", cfg.Server.Port))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), handler.Routes(h)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Println("Application started successfully")
}