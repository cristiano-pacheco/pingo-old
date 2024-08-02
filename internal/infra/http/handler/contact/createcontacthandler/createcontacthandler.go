// Package createcontacthandler provides a handler to create a contact.
package createcontacthandler

import (
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/contact/createcontactuc"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/response"
)

type Handler struct {
	createContactUseCase *createcontactuc.UseCase
}

func New(useCase *createcontactuc.UseCase) *Handler {
	return &Handler{createContactUseCase: useCase}
}

func (h *Handler) Execute(w http.ResponseWriter, r *http.Request) {
	var in input
	err := request.ReadJSON(w, r, &in)
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	vr := validateInput(in)
	if !vr.IsValid {
		response.ValidationFailedResponse(w, r, vr)
		return
	}

	useCaseInput := mapToUseCaseInput(in, r)
	useCaseOutput, err := h.createContactUseCase.Execute(useCaseInput)
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
