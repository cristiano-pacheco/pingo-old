// Package findcontactlistuc contains the use case to find a contact list.
package findcontactlistuc

import (
	"github.com/cristiano-pacheco/pingo/internal/application/repository/contactrepo"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/identitydm"
)

type UseCase struct {
	contactRepo contactrepo.ContactRepository
}

func New(contactRepo contactrepo.ContactRepository) *UseCase {
	return &UseCase{contactRepo: contactRepo}
}

func (uc *UseCase) Execute(input *Input) (*Output, error) {
	userIDVo, err := identitydm.Restore(input.UserID)
	if err != nil {
		return nil, err
	}

	contactList, err := uc.contactRepo.FindListByUserID(*userIDVo)
	if err != nil {
		return nil, err
	}

	output := mapContactListToOutput(contactList)

	return output, nil
}
