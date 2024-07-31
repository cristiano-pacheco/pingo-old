// Package activateuserhandler provides a handler to activate the user account.
package activateuserhandler

import (
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/user/activateuseruc"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/response"
)

type Handler struct {
	activateUserUseCase *activateuseruc.UseCase
}

func New(useCase *activateuseruc.UseCase) *Handler {
	return &Handler{activateUserUseCase: useCase}
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

	useCaseInput, err := mapInputToUseCaseInput(&in)
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	err = h.activateUserUseCase.Execute(useCaseInput)
	if err != nil {
		response.ServerErrorResponse(w, r, err)
		return
	}

	response.EmptyOKResponse(w)
}
