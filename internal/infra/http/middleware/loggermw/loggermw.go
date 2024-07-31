// Package loggermw contains a middeware to add a logger to the request context.
package loggermw

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
)

func AddLoggerToContextMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), request.LoggerContextKey, logger)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
