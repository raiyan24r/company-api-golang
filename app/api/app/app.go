package app

import (
	"context"

	"go.uber.org/zap"
)

type App struct {
	Config Config
	Log    *zap.Logger
}

func New(ctx context.Context, cfg Config, logger *zap.Logger) *App {
	return &App{
		Config: cfg,
		Log:    logger,
	}
}
