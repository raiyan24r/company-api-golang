package main

import (
	"company-api/app/api/app"
	"company-api/foundation/logger"
	"log"
)

func main() {
	log.Println("Starting the application...")

	_, err := logger.New()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	_, err = app.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}


}


