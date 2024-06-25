package createuseruc

import (
	"fmt"

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

	dbUser, _ := uc.userRepo.FindByEmail(user.Email)
	if dbUser != nil {
		return nil, fmt.Errorf("the email %s is already in use", input.Email)
	}

	err = uc.userRepo.Create(*user)
	if err != nil {
		return nil, err
	}

	output := uc.mapper.mapUserToOutput(*user)

	return output, nil
}
