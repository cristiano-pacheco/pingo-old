package updatecontactuc

import (
	"time"

	"github.com/cristiano-pacheco/pingo/internal/domain/model/contactdm"
)

func mapInputToContact(in *Input) (*contactdm.Contact, error) {
	now := time.Now().UTC()
	contact, err := contactdm.Restore(
		in.ID, in.UserID, in.Name, in.ContactType, in.ContactValue, in.IsEnabled, now, now,
	)
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
