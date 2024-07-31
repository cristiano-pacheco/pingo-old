// Package updateuserhandler provides a handler to update the user.
package updateuserhandler

import (
	"fmt"
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/user/updateuseruc"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/response"
)

type Handler struct {
	updateUserUseCase *updateuseruc.UseCase
}

func New(useCase *updateuseruc.UseCase) *Handler {
	return &Handler{updateUserUseCase: useCase}
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

	useCaseInput := &updateuseruc.Input{
		UserID: userID,
		Name:   in.Name,
	}

	err = h.updateUserUseCase.Execute(useCaseInput)
	if err != nil {
		response.ServerErrorResponse(w, r, err)
		return
	}

	response.NoContentResponse(w)
}
