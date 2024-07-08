// Package hashds contains a domain service to handle hash.
package hashds

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

const defaultTotalRandomBytesSize = 128

func (s *Service) GenerateFromPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func (s *Service) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func (s *Service) GenerateRandomBytes() ([]byte, error) {
	buffer := make([]byte, defaultTotalRandomBytesSize)
	_, err := rand.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
