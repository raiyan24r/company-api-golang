package handler

import "go.uber.org/zap"

type Handler struct {
	Log *zap.Logger
}

func New(logger *zap.Logger) Handler {
	return Handler{
		Log: logger,
	}
}