// Package hashds contains a domain service to handle hash.
package hashds

import (
	"golang.org/x/crypto/bcrypt"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) GenerateFromPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func (s *Service) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
