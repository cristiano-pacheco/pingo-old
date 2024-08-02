package updatecontacthandler

import (
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/contact/updatecontactuc"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
)

func mapToUseCaseInput(in input, r *http.Request) *updatecontactuc.Input {
	input := updatecontactuc.Input{
		Name:         in.Name,
		ContactType:  in.ContactType,
		ContactValue: in.ContactValue,
		IsEnabled:    in.IsEnabled,
	}
	input.UserID = request.GetUserIDFromContext(r)
	input.ID = request.GetParam(r, "contactId")
	return &input
}
