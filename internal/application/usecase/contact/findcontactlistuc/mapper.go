package findcontactlistuc

import (
	"github.com/cristiano-pacheco/pingo/internal/domain/model/contactdm"
)

func mapContactListToOutput(contacts []*contactdm.Contact) *Output {
	var output Output
	for _, c := range contacts {
		contact := &contact{
			ID:           c.ID.String(),
			UserID:       c.UserID.String(),
			Name:         c.Name.String(),
			ContactType:  c.ContactType(),
			ContactValue: c.ContactValue(),
		}
		output.Items = append(output.Items, contact)
	}
	return &output
}
