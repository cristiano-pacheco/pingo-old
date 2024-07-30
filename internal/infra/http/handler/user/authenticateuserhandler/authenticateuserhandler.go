// Package authenticateuserhandler provides a handler to authenticate the user.
package authenticateuserhandler

import (
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/user/authenticateuseruc"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/response"
)

type Handler struct {
	usecase authenticateuseruc.UseCase
}

func New(usecase *authenticateuseruc.UseCase) *Handler {
	return &Handler{usecase: *usecase}
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

	useCaseInput := &authenticateuseruc.Input{Email: in.Email, Password: in.Password}
	useCaseOutput, err := h.usecase.Execute(useCaseInput)

	// handle invalid credentials and return the http status 401 instead
	if err != nil {
		response.ServerErrorResponse(w, r, err)
		return
	}

	out := &output{Token: useCaseOutput.Token}
	envelope := &response.Envelope{"data": out}
	err = response.JSONResponse(w, http.StatusOK, *envelope, nil)
	if err != nil {
		response.ServerErrorResponse(w, r, err)
		return
	}
}
