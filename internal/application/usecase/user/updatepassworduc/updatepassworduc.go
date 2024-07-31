// Package updatepassworduc contains the use case to update the user password.
package updatepassworduc

import (
	"github.com/cristiano-pacheco/pingo/internal/application/repository/userrepo"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/identitydm"
	"github.com/cristiano-pacheco/pingo/internal/domain/service/hashds"
)

type UseCase struct {
	userRepo    userrepo.UserRepository
	hashService hashds.Service
}

func New(
	userRepo userrepo.UserRepository,
	hashService hashds.Service,
) *UseCase {
	return &UseCase{
		userRepo:    userRepo,
		hashService: hashService,
	}
}

func (uc *UseCase) Execute(input *Input) error {
	// finds the user by ID
	idVo, err := identitydm.Restore(input.UserID)
	if err != nil {
		return err
	}

	user, err := uc.userRepo.FindByID(*idVo)
	if err != nil {
		return err
	}

	// verifies if the current password is right
	err = uc.hashService.CompareHashAndPassword(user.PasswordHash, []byte(input.CurrentPassword))
	if err != nil {
		return err
	}

	// creates a password hash
	passwordHash, err := uc.hashService.GenerateFromPassword([]byte(input.NewPassword))
	if err != nil {
		return err
	}

	// updates the password hash in the database
	user.PasswordHash = passwordHash
	err = uc.userRepo.UpdatePassword(*user)
	if err != nil {
		return err
	}

	return nil
}
