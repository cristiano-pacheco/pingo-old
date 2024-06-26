// Package createuseruc contains the use case to create the user.
package createuseruc

import (
	"github.com/cristiano-pacheco/pingo/internal/application/repository/userrepo"
)

type UseCase struct {
	userRepo userrepo.UserRepository
	mapper   Mapper
}

func New(userRepo userrepo.UserRepository, mapper Mapper) *UseCase {
	return &UseCase{
		userRepo: userRepo,
		mapper:   mapper,
	}
}

func (uc *UseCase) Execute(input *Input) (*Output, error) {
	user, err := uc.mapper.mapInputToNewUser(input)
	if err != nil {
		return nil, err
	}

	err = uc.userRepo.Create(*user)
	if err != nil {
		return nil, err
	}

	output := uc.mapper.mapUserToOutput(*user)

	return output, nil
}
