package middlewares

import (
	"log/slog"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			slog.Info(
				"request logger middleware",
				"method", r.Method,
				"proto", r.Proto,
				"path", r.URL.Path,
				"query", r.URL.Query(),
				"ip", r.RemoteAddr,
				"user_agent", r.UserAgent(),
				"request_id", GetReqestID(r.Context()),
				"elapsed", time.Since(start).String(),
			)
		}()
		next.ServeHTTP(w, r)
	})
}
