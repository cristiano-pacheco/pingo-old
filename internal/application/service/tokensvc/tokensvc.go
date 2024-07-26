package tokensvc

import "github.com/cristiano-pacheco/pingo/internal/domain/model/userdm"

type TokenService interface {
	GenerateToken(user *userdm.User) (string, error)
	ValidateToken(token string) error
}
