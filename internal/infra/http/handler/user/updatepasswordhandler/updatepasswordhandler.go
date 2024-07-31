// Package updatepasswordhandler provides a handler to update the user password.
package updatepasswordhandler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/user/updatepassworduc"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/response"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	updatePasswordUseCase *updatepassworduc.UseCase
}

func New(useCase *updatepassworduc.UseCase) *Handler {
	return &Handler{updatePasswordUseCase: useCase}
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

	userID := r.Context().Value(request.UserIDContextKey).(string)
	if userID == "" {
		response.BadRequestResponse(w, r, fmt.Errorf("the user id is invalid"))
		return
	}

	useCaseInput := &updatepassworduc.Input{
		UserID:          userID,
		CurrentPassword: in.CurrentPassword,
		NewPassword:     in.NewPassword,
	}

	err = h.updatePasswordUseCase.Execute(useCaseInput)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			data := &response.Envelope{"error": "the current password does not match"}
			response.JSONResponse(w, http.StatusUnprocessableEntity, *data, nil)
			return
		}
		response.ServerErrorResponse(w, r, err)
		return
	}

	response.EmptyOKResponse(w)
}
