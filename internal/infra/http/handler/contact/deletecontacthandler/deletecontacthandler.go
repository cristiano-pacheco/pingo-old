// Package deletecontacthandler provides a handler to delete a contact.
package deletecontacthandler

import (
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/contact/deletecontactuc"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/response"
)

type Handler struct {
	deleteContactUseCase *deletecontactuc.UseCase
}

func New(useCase *deletecontactuc.UseCase) *Handler {
	return &Handler{deleteContactUseCase: useCase}
}

func (h *Handler) Execute(w http.ResponseWriter, r *http.Request) {
	useCaseInput := deletecontactuc.Input{
		ID:     request.GetParam(r, "contactId"),
		UserID: request.GetUserIDFromContext(r),
	}

	err := h.deleteContactUseCase.Execute(&useCaseInput)
	if err != nil {
		response.ServerErrorResponse(w, r, err)
		return
	}

	response.EmptyOKResponse(w)
}
