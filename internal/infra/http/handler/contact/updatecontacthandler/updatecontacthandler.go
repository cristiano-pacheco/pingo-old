// Package updatecontacthandler provides a handler to update a contact.
package updatecontacthandler

import (
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/contact/updatecontactuc"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/response"
)

type Handler struct {
	updateContactUseCase *updatecontactuc.UseCase
}

func New(useCase *updatecontactuc.UseCase) *Handler {
	return &Handler{updateContactUseCase: useCase}
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

	useCaseInput := mapToUseCaseInput(in, r)

	_, err = h.updateContactUseCase.Execute(useCaseInput)
	if err != nil {
		response.ServerErrorResponse(w, r, err)
		return
	}

	response.EmptyOKResponse(w)
}
