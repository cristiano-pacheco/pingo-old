package contactrepo

import (
	"github.com/cristiano-pacheco/pingo/internal/domain/model/contactdm"
)

func mapContactDBToContact(contactdb *ContactDB) (*contactdm.Contact, error) {
	contact, err := contactdm.Restore(
		contactdb.ID,
		contactdb.UserID,
		contactdb.Name,
		contactdb.ContactType,
		contactdb.ContactData,
		contactdb.IsEnabled,
		contactdb.CreatedAT,
		contactdb.UpdatedAT,
	)

	if err != nil {
		return nil, err
	}

	return contact, nil
}
