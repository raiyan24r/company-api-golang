package app

import (
	"company-api/business/database"
	"context"

	"go.uber.org/zap"
)

type App struct {
	Config Config
	Log    *zap.Logger
	DbRepo *database.Database
}

func New(ctx context.Context, cfg Config, logger *zap.Logger, dbRepo *database.Database) *App {
	return &App{
		Config: cfg,
		Log:    logger,
		DbRepo: dbRepo,
	}
}
