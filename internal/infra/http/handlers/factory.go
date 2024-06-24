package handlers

import "github.com/cristiano-pacheco/pingo/internal/infra/http/handlers/pinghandler"

type Handlers struct {
	PingHandler *pinghandler.Handler
}

func New() *Handlers {
	ph := pinghandler.New()
	return &Handlers{
		PingHandler: ph,
	}
}
