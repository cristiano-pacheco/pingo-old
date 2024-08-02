package findcontactlisthandler

import (
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/contact/findcontactlistuc"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
)

func mapToUseCaseInput(r *http.Request) *findcontactlistuc.Input {
	input := findcontactlistuc.Input{
		UserID: request.GetUserIdFromContext(r),
	}
	return &input
}

func mapToOutput(out *findcontactlistuc.Output) *output {
	var o output
	for _, c := range out.Items {
		c := &contact{
			ID:           c.ID,
			Name:         c.Name,
			ContactType:  c.ContactType,
			ContactValue: c.ContactValue,
			IsEnabled:    c.IsEnabled,
		}
		o.Items = append(o.Items, c)
	}
	return &o
}
