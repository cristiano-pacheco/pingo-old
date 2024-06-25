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
	var i input
	err := request.ReadJSON(w, r, i)
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	useCaseInput := &createuseruc.Input{
		Name:     i.Name,
		Email:    i.Email,
		Password: i.Password,
	}

	_, err = h.createUserUseCase.Execute(useCaseInput)
	if err != nil {
		response.ServerErrorResponse(w, r, err)
		return
	}

	err = response.JSONResponse(w, http.StatusCreated, nil, nil)
	if err != nil {
		response.ServerErrorResponse(w, r, err)
		return
	}
}
