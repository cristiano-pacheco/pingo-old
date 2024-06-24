package router

import (
	"log/slog"

	"github.com/cristiano-pacheco/pingo/internal/infra/http/handlers"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	Mux *chi.Mux
}

func New(handlers *handlers.Handlers, logger *slog.Logger) *Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.NotFound(response.NotFoundResponse)
	router.MethodNotAllowed(response.MethodNotAllowedResponse)

	router.Get("/v1/ping", handlers.PingHandler.Execute)

	return &Router{Mux: router}
}
