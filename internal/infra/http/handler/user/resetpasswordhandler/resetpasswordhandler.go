// Package resetpasswordhandler provides a handler to process the user reset password.
package resetpasswordhandler

import (
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/user/resetpassworduc"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/response"
)

type Handler struct {
	resetPasswordUseCase *resetpassworduc.UseCase
}

func New(useCase *resetpassworduc.UseCase) *Handler {
	return &Handler{resetPasswordUseCase: useCase}
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

	err = h.resetPasswordUseCase.Execute(useCaseInput)
	if err != nil {
		response.ServerErrorResponse(w, r, err)
		return
	}

	response.EmptyOKResponse(w)
}
