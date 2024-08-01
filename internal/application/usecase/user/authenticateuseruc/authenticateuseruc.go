// Package authenticateuseruc contains the use case to authenticate the user.
package authenticateuseruc

import (
	"github.com/cristiano-pacheco/pingo/internal/application/repository/userrepo"
	"github.com/cristiano-pacheco/pingo/internal/application/service/tokensvc"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/emaildm"
	"github.com/cristiano-pacheco/pingo/internal/domain/service/hashds"
)

type UseCase struct {
	tokenService   tokensvc.TokenService
	userRepository userrepo.UserRepository
	hashService    hashds.Service
}

func New(
	tokenService tokensvc.TokenService,
	userRepo userrepo.UserRepository,
	hashService hashds.Service,
) *UseCase {
	return &UseCase{
		tokenService:   tokenService,
		userRepository: userRepo,
		hashService:    hashService,
	}
}

func (uc *UseCase) Execute(in *Input) (*Output, error) {
	emailVo, err := emaildm.New(in.Email)
	if err != nil {
		return nil, err
	}

	user, err := uc.userRepository.FindByEmail(*emailVo)
	if err != nil {
		return nil, err
	}

	err = uc.hashService.CompareHashAndPassword(user.PasswordHash, []byte(in.Password))
	if err != nil {
		return nil, err
	}

	jwtToken, err := uc.tokenService.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	output := &Output{Token: jwtToken}
	return output, nil
}
