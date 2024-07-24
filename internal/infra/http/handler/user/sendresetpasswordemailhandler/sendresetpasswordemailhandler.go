// Package sendresetpasswordemailhandler handles the user reset password request.
package sendresetpasswordemailhandler

import (
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/user/sendresetpasswordemailuc"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/response"
)

type Handler struct {
	resetPasswordUseCase *sendresetpasswordemailuc.UseCase
}

func New(useCase *sendresetpasswordemailuc.UseCase) *Handler {
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

	useCaseInput := &sendresetpasswordemailuc.Input{Email: in.Email}

	err = h.resetPasswordUseCase.Execute(useCaseInput)
	if err != nil {
		response.ServerErrorResponse(w, r, err)
		return
	}

	response.EmptyResponse(w)
}
