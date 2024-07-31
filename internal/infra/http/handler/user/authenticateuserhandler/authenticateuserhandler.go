// Package authenticateuserhandler provides a handler to authenticate the user.
package authenticateuserhandler

import (
	"errors"
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/user/authenticateuseruc"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/response"
	"golang.org/x/crypto/bcrypt"
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

	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			data := &response.Envelope{"error": "invalid credentials"}
			response.JSONResponse(w, http.StatusUnauthorized, *data, nil)
			return
		}
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
