// Package keydm contains the value object to carry the hash key values.
package keydm

import (
	"crypto/rsa"
	"errors"
)

var (
	ErrKeyMustBePEMEncoded = errors.New("invalid key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	ErrNotRSAPrivateKey    = errors.New("key is not a valid RSA private key")
)

type Key struct {
	privatePEM string
	publicPEM  string
	privateRSA *rsa.PrivateKey
}

func New(keyBytes []byte) (*Key, error) {
	privateRSA, err := mapPEMToRSAPrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}

	privatePEM := string(keyBytes)

	publicPEM, err := mapRSAPrivateKeyToPublicKey(privateRSA)
	if err != nil {
		return nil, err
	}

	key := &Key{privatePEM: privatePEM, publicPEM: publicPEM, privateRSA: privateRSA}

	return key, nil
}

func (k *Key) PrivateRSA() *rsa.PrivateKey {
	return k.privateRSA
}

func (k *Key) PrivateKey() string {
	return k.privatePEM
}

func (k *Key) PublicKey() string {
	return k.publicPEM
}
