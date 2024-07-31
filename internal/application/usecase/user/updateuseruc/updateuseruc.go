// Package updateuseruc contains the use case to update the user.
package updateuseruc

import (
	"github.com/cristiano-pacheco/pingo/internal/application/repository/userrepo"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/identitydm"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/userdm"
)

type UseCase struct {
	userRepo userrepo.UserRepository
}

func New(
	userRepo userrepo.UserRepository,
) *UseCase {
	return &UseCase{
		userRepo: userRepo,
	}
}

func (uc *UseCase) Execute(input *Input) error {
	idVo, err := identitydm.Restore(input.UserID)
	if err != nil {
		return err
	}

	user, err := uc.userRepo.FindByID(*idVo)
	if err != nil {
		return err
	}

	nameVo, err := userdm.NewName(input.Name)
	if err != nil {
		return err
	}

	user.Name = *nameVo
	err = uc.userRepo.Update(*user)
	if err != nil {
		return err
	}

	return nil
}
