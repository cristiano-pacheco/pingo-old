// Package resetpassworduc contains the use case to process the user reset password.
package resetpassworduc

import (
	"fmt"

	"github.com/cristiano-pacheco/pingo/internal/application/repository/userrepo"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/identitydm"
	"github.com/cristiano-pacheco/pingo/internal/domain/service/hashds"
)

type UseCase struct {
	userRepo    userrepo.UserRepository
	hashService *hashds.Service
}

func New(
	userRepo userrepo.UserRepository,
	hashService *hashds.Service,
) *UseCase {
	return &UseCase{
		userRepo:    userRepo,
		hashService: hashService,
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

	resetPasswordToken := string(user.ResetPasswordToken)
	if input.ResetPasswordToken != resetPasswordToken {
		return fmt.Errorf("the reset password token is invalid")
	}

	passwordHash, err := uc.hashService.GenerateFromPassword([]byte(input.Password))
	if err != nil {
		return err
	}

	user.PasswordHash = passwordHash
	err = uc.userRepo.UpdatePassword(*user)
	if err != nil {
		return err
	}

	return nil
}
