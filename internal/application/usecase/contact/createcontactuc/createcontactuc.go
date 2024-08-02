// Package createcontactuc contains the use case to create a contact.
package createcontactuc

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

	err = uc.contactRepo.Create(*contact)
	if err != nil {
		return nil, err
	}

	output := mapContactToOutput(contact)

	return output, nil
}
