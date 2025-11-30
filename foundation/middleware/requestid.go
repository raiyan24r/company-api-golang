package middleware

import (
	"context"
	"net/http"
)

func RequestID() Middleware {
	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			reqID := r.Header.Get("X-Request-ID")
			if reqID == "" {
				reqID = "generated-request-id" // In real implementation, generate a unique ID
			}
			ctx = context.WithValue(ctx, "requestID", reqID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}