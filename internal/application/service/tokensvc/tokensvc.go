// Package tokensvc contains the TokenService contract.
package tokensvc

import (
	"context"

	"github.com/cristiano-pacheco/pingo/internal/domain/model/authdm"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/userdm"
)

type TokenService interface {
	GenerateToken(user *userdm.User) (string, error)
	ParseToken(ctx context.Context, bearerToken string) (*authdm.JWTClaims, error)
}
