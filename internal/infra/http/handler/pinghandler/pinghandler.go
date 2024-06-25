package pinghandler

import (
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/infra/http/response"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Execute(w http.ResponseWriter, r *http.Request) {
	data := &response.Envelope{"data": "Pong!"}
	err := response.JSONResponse(w, http.StatusOK, *data, nil)
	if err != nil {
		response.ServerErrorResponse(w, r, err)
	}
}
