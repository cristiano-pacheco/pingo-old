// Package tokensvc contains the service that handles the authentication.
package tokensvc

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cristiano-pacheco/pingo/internal/application/repository/userrepo"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/authdm"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/identitydm"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/keydm"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/userdm"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	userRepo userrepo.UserRepository
	key      *keydm.Key
	parser   *jwt.Parser
	issuer   string
}

func New(
	userRepo userrepo.UserRepository,
	key *keydm.Key,
	parser *jwt.Parser,
	iss string,
) *TokenService {
	return &TokenService{userRepo: userRepo, key: key, parser: parser, issuer: iss}
}

func (s *TokenService) GenerateToken(user *userdm.User) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    s.issuer,
		Subject:   user.ID.String(),
	}

	method := jwt.GetSigningMethod(jwt.SigningMethodRS256.Name)
	token := jwt.NewWithClaims(method, claims)

	signedToken, err := token.SignedString(s.key.PrivateRSA())
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ParseToken processes the token to validate the sender's token is valid.
func (s *TokenService) ParseToken(ctx context.Context, bearerToken string) (*authdm.JWTClaims, error) {
	if !strings.HasPrefix(bearerToken, "Bearer ") {
		return nil, errors.New("expected authorization header format: Bearer <token>")
	}

	jwt := strings.TrimSpace(bearerToken[7:])

	var claims authdm.JWTClaims
	_, _, err := s.parser.ParseUnverified(jwt, &claims)
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	input := map[string]any{
		"key":   s.key.PublicKey(),
		"token": jwt,
	}

	err = validateOpaPolicy(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("authentication failed : %w", err)
	}

	// Check the database for this user to verify they are still enabled.
	if err := s.isUserActivated(&claims); err != nil {
		return nil, fmt.Errorf("user not enabled : %w", err)
	}

	return &claims, nil
}

// isUserEnabled hits the database and checks the user is not disabled.
func (s *TokenService) isUserActivated(claims *authdm.JWTClaims) error {
	userID, err := identitydm.Restore(claims.Subject)
	if err != nil {
		return fmt.Errorf("parse user: %w", err)
	}

	user, err := s.userRepo.FindByID(*userID)
	if err != nil {
		return fmt.Errorf("query user: %w", err)
	}

	if !user.IsActivated() {
		return fmt.Errorf("user disabled")
	}

	return nil
}
