package privatekeydm

import (
	"crypto/rsa"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type PrivateKey struct {
	value *rsa.PrivateKey
}

func New(pemData []byte) (*PrivateKey, error) {
	pk, err := jwt.ParseRSAPrivateKeyFromPEM(pemData)
	if err != nil {
		return nil, fmt.Errorf("parsing auth private key: %w", err)
	}

	return &PrivateKey{value: pk}, nil
}

func (pk *PrivateKey) Value() *rsa.PrivateKey {
	return pk.value
}
