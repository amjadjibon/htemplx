package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"htemplx/pkg/logger"
)

func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		ctx = context.WithValue(ctx, logger.RequestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if reqID, ok := ctx.Value(logger.RequestIDKey).(string); ok {
		return reqID
	}

	return ""
}
