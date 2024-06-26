// Package createuserhandler provides a handler to handle the user creation.
package createuserhandler

import (
	"fmt"
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/user/createuseruc"
	"github.com/cristiano-pacheco/pingo/internal/infra/database/dberror"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/response"
	"github.com/lib/pq"
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

	vr := validateInput(in)
	if !vr.IsValid {
		response.ValidationFailedResponse(w, r, vr)
		return
	}

	useCaseInput := mapInputToUseCaseInput(in)

	useCaseOutput, err := h.createUserUseCase.Execute(useCaseInput)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == dberror.ErrUniqueViolationCode {
				newErr := fmt.Errorf("the email %s is already in use", in.Email)
				response.BadRequestResponse(w, r, newErr)
				return
			}
		}
		response.ServerErrorResponse(w, r, err)
		return
	}

	out := mapUseCaseOutputToOutput(useCaseOutput)

	envelope := &response.Envelope{"data": out}
	err = response.JSONResponse(w, http.StatusCreated, *envelope, nil)
	if err != nil {
		response.ServerErrorResponse(w, r, err)
		return
	}
}
