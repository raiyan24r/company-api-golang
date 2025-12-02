# ---- Migration settings ----
MIGRATE_PATH := business/database/migration
# Build DSN from config.yml (requires yq or parse manually)
DB_HOST := $(shell grep -A6 "^db:" config.yml | grep "host:" | awk '{print $$2}' | tr -d '"')
DB_PORT := $(shell grep -A6 "^db:" config.yml | grep "port:" | awk '{print $$2}')
DB_USER := $(shell grep -A6 "^db:" config.yml | grep "user:" | awk '{print $$2}' | tr -d '"')
DB_PASS := $(shell grep -A6 "^db:" config.yml | grep "password:" | awk '{print $$2}' | tr -d '"')
DB_NAME := $(shell grep -A6 "^db:" config.yml | grep "name:" | awk '{print $$2}' | tr -d '"')
DB_DSN ?= mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)


deps:
	go mod tidy
	go mod vendor

start:
	go run ./app/api

docker-up:
	docker-compose up -d
docker-down:
	docker-compose down
	@echo "Cleaning up stale processes on project ports..."
	@sudo lsof -ti :8080,:8081,:3307 2>/dev/null | xargs -r sudo kill -9 2>/dev/null || true

migrate-up:
	migrate -path $(MIGRATE_PATH) -database "$(DB_DSN)" up

migrate-down:
	migrate -path $(MIGRATE_PATH) -database "$(DB_DSN)" down 1

# Create a new migration (use NAME=<snake_case_name>)
migrate-create:
	@if [ -z "$(NAME)" ]; then echo "Usage: make migrate-create NAME=<snake_case_name>"; exit 1; fi
	migrate create -ext sql -dir $(MIGRATE_PATH) -seq $(NAME)