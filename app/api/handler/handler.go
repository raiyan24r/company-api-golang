package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Handler struct {
	Log *zap.Logger
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func New(logger *zap.Logger) Handler {
	return Handler{
		Log: logger,
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details"`
	Service string `json:"service"`
	Type    string `json:"type"`
	Reason  string `json:"reason"`
	Time    string `json:"time,"`
}

func (h Handler) ErrorWrapper(mainFunc HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := mainFunc(w, r)
		if err != nil {
			h.Log.Error("Handler error", zap.Error(err))
			h.writeErrorResponse(r.Context(), w, http.StatusInternalServerError, err.Error())
		}
	}
}

func (h Handler) writeResponse(ctx context.Context, w http.ResponseWriter, statusCode int, data any) error {
	w.WriteHeader(statusCode)

	if statusCode == http.StatusNoContent || data == nil {
		return nil
	}

	if b, ok := data.([]byte); ok {
		if _, err := w.Write(b); err != nil {
			h.Log.Error("Failed to write response", zap.Error(err))
			return err
		}
		return nil
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.Log.Error("Failed to send response", zap.Error(err))
	}

	return nil
}

func (h Handler) writeErrorResponse(ctx context.Context, w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)

	errorResp := ErrorResponse{
		Message: message,
		Service: "company-api",
		Type:    "error",
		Time:    time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
	}

	if err := json.NewEncoder(w).Encode(errorResp); err != nil {
		h.Log.Error("Failed to send error response", zap.Error(err))
	}
}
