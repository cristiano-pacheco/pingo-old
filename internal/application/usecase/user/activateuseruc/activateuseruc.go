// Package activateuseruc contains the use case to activate the user account.
package activateuseruc

import (
	"fmt"

	"github.com/cristiano-pacheco/pingo/internal/application/repository/userrepo"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/identitydm"
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
	ID, err := identitydm.Restore(input.ID)
	if err != nil {
		return err
	}

	user, err := uc.userRepo.FindByID(*ID)
	if err != nil {
		return err
	}

	accountConfToken := string(user.AccountConfirmationToken)
	if input.AccountConfToken != accountConfToken {
		return fmt.Errorf("the account confirmation token is invalid")
	}

	err = uc.userRepo.ActivateAccount(*user)
	if err != nil {
		return err
	}

	return nil
}
