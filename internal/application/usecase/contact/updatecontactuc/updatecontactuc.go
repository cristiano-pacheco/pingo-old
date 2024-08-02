// Package updatecontactuc contains the use case to update a contact.
package updatecontactuc

import (
	"github.com/cristiano-pacheco/pingo/internal/application/repository/contactrepo"
)

type UseCase struct {
	contactRepo contactrepo.ContactRepository
}

func New(
	contactRepo contactrepo.ContactRepository,
) *UseCase {
	return &UseCase{
		contactRepo: contactRepo,
	}
}

func (uc *UseCase) Execute(input *Input) (*Output, error) {
	contact, err := mapInputToContact(input)
	if err != nil {
		return nil, err
	}

	// It will check if the contact belongs to the user
	_, err = uc.contactRepo.FindByIDAndUserID(contact.ID, contact.UserID)
	if err != nil {
		return nil, err
	}

	err = uc.contactRepo.Update(*contact)
	if err != nil {
		return nil, err
	}

	output := mapContactToOutput(contact)

	return output, nil
}
