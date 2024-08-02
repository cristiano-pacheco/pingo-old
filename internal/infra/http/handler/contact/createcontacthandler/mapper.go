package createcontacthandler

import (
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/contact/createcontactuc"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
)

func mapToUseCaseInput(in input, r *http.Request) *createcontactuc.Input {
	input := createcontactuc.Input{
		Name:         in.Name,
		ContactType:  in.ContactType,
		ContactValue: in.ContactValue,
		IsEnabled:    in.IsEnabled,
	}
	input.UserID = request.GetUserIdFromContext(r)
	return &input
}

func mapToOutput(out *createcontactuc.Output) *output {
	return &output{
		ID:           out.ID,
		Name:         out.Name,
		ContactType:  out.ContactType,
		ContactValue: out.ContactValue,
		IsEnabled:    out.IsEnabled,
	}
}
