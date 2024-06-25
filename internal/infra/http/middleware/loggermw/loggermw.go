package loggermw

import (
	"context"
	"log/slog"
	"net/http"
)

type contextKey string

const LOGGER_CONTEXT_KEY = "logger"

func AddLoggerToContextMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), contextKey(LOGGER_CONTEXT_KEY), logger)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
