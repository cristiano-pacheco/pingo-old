// Package findcontactlisthandler provides a handler to retreive a list of contacts by user id.
package findcontactlisthandler

import (
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/contact/findcontactlistuc"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/response"
)

type Handler struct {
	findContactListUseCase *findcontactlistuc.UseCase
}

func New(useCase *findcontactlistuc.UseCase) *Handler {
	return &Handler{findContactListUseCase: useCase}
}

func (h *Handler) Execute(w http.ResponseWriter, r *http.Request) {
	useCaseInput := mapToUseCaseInput(r)

	useCaseOutput, err := h.findContactListUseCase.Execute(useCaseInput)
	if err != nil {
		response.ServerErrorResponse(w, r, err)
		return
	}

	out := mapToOutput(useCaseOutput)

	envelope := &response.Envelope{"data": out}
	err = response.JSONResponse(w, http.StatusCreated, *envelope, nil)
	if err != nil {
		response.ServerErrorResponse(w, r, err)
		return
	}
}
