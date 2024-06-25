// Package createuserhandler provides a handler to handle the user creation.
package createuserhandler

import (
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/user/createuseruc"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/response"
)

type Handler struct {
	createUserUseCase *createuseruc.UseCase
}

func New(useCase *createuseruc.UseCase) *Handler {
	return &Handler{createUserUseCase: useCase}
}

func (h *Handler) Execute(w http.ResponseWriter, r *http.Request) {
	var in input
	err := request.ReadJSON(w, r, &in)
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	useCaseInput := &createuseruc.Input{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}

	useCaseOutput, err := h.createUserUseCase.Execute(useCaseInput)
	if err != nil {
		response.ServerErrorResponse(w, r, err)
		return
	}

	out := &output{
		ID:    useCaseOutput.ID,
		Name:  useCaseOutput.Name,
		Email: useCaseOutput.Email,
	}

	envelope := &response.Envelope{"data": out}
	err = response.JSONResponse(w, http.StatusCreated, *envelope, nil)
	if err != nil {
		response.ServerErrorResponse(w, r, err)
		return
	}
}
