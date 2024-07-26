package tokensvc

import (
	"time"

	"github.com/cristiano-pacheco/pingo/internal/domain/model/privatekeydm"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/userdm"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	privateKey *privatekeydm.PrivateKey
	issuerName string
}

func New(privateKey *privatekeydm.PrivateKey, issuerName string) *TokenService {
	return &TokenService{privateKey: privateKey, issuerName: issuerName}
}

func (s *TokenService) GenerateToken(user *userdm.User) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    s.issuerName,
		Subject:   user.ID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	result, err := token.SignedString(s.privateKey.Value())
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *TokenService) ValidateToken(token string) error {
	return nil
}
