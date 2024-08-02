package deletecontactuc

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

func (uc *UseCase) Execute(input *Input) error {
	idVo, err := identitydm.Restore(input.ID)
	if err != nil {
		return err
	}

	userIDVo, err := identitydm.Restore(input.UserID)
	if err != nil {
		return err
	}

	contact, err := uc.contactRepo.FindByIDAndUserID(*idVo, *userIDVo)
	if err != nil {
		return err
	}

	err = uc.contactRepo.Delete(*contact)
	if err != nil {
		return err
	}

	return nil
}
