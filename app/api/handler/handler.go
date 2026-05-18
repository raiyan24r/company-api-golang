package handler

import (
	"company-api/app/api/app"
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Handler struct {
	App      *app.App
	Validate *validator.Validate
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func New(a *app.App) Handler {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return Handler{
		App:      a,
		Validate: v,
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

type ValidationErrorResponse struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func (h Handler) ErrorWrapper(mainFunc HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := mainFunc(w, r)
		if err != nil {
			h.App.Log.Error("Handler error", zap.Error(err))
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
			h.App.Log.Error("Failed to write response", zap.Error(err))
			return err
		}
		return nil
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.App.Log.Error("Failed to send response", zap.Error(err))
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
		h.App.Log.Error("Failed to send error response", zap.Error(err))
	}
}

func (h Handler) writeValidationErrorResponse(ctx context.Context, w http.ResponseWriter, err error) {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		h.writeErrorResponse(ctx, w, http.StatusBadRequest, "invalid request payload")
		return
	}

	errors := make(map[string]string)
	for _, fieldErr := range validationErrors {
		field := fieldErr.Field()
		errors[field] = validationMessage(fieldErr)
	}

	_ = h.writeResponse(ctx, w, http.StatusBadRequest, ValidationErrorResponse{
		Message: "validation failed",
		Errors:  errors,
	})
}

func validationMessage(fieldErr validator.FieldError) string {
	field := fieldErr.Field()

	switch fieldErr.Tag() {
	case "required":
		return field + " is required"
	case "min":
		return field + " must be at least " + fieldErr.Param() + " characters"
	case "max":
		return field + " must be at most " + fieldErr.Param() + " characters"
	case "gte":
		return field + " must be greater than or equal to " + fieldErr.Param()
	default:
		return field + " is invalid"
	}
}
