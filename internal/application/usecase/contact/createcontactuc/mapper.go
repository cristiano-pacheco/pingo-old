package createcontactuc

import (
	"github.com/cristiano-pacheco/pingo/internal/domain/model/contactdm"
)

func mapInputToContact(input *Input) (*contactdm.Contact, error) {
	contact, err := contactdm.New(input.UserID, input.Name, input.ContactType, input.ContactValue, input.IsEnabled)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func mapContactToOutput(c *contactdm.Contact) *Output {
	output := &Output{
		ID:           c.ID.String(),
		UserID:       c.UserID.String(),
		Name:         c.Name.String(),
		ContactType:  c.ContactType(),
		ContactValue: c.ContactValue(),
	}
	return output
}
