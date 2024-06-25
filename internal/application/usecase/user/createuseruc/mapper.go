package createuseruc

import (
	"github.com/cristiano-pacheco/pingo/internal/domain/model/userdm"
	"github.com/cristiano-pacheco/pingo/internal/domain/service/hashds"
)

type Mapper struct {
	hashService *hashds.Service
}

func NewMapper(hs *hashds.Service) *Mapper {
	return &Mapper{hashService: hs}
}

func (m *Mapper) mapInputToNewUser(input Input) (*userdm.User, error) {
	passwordHash, err := m.hashService.GenerateFromPassword([]byte(input.Password))
	if err != nil {
		return nil, err
	}
	user, err := userdm.NewUser(input.Name, input.Email, passwordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *Mapper) mapUserToOutput(user userdm.User) *Output {
	return &Output{
		ID:    user.ID.String(),
		Name:  user.Name.String(),
		Email: user.Email.String(),
	}
}
