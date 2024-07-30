package keydm

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func mapPEMToRSAPrivateKey(key []byte) (*rsa.PrivateKey, error) {
	var err error

	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, ErrKeyMustBePEMEncoded
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
			return nil, err
		}
	}

	var pkey *rsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PrivateKey); !ok {
		return nil, ErrNotRSAPrivateKey
	}

	return pkey, nil
}

func mapRSAPrivateKeyToPublicKey(privateRSA *rsa.PrivateKey) (string, error) {
	asn1Bytes, err := x509.MarshalPKIXPublicKey(&privateRSA.PublicKey)
	if err != nil {
		return "", fmt.Errorf("marshaling public key: %w", err)
	}

	publicBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	var buf bytes.Buffer
	if err := pem.Encode(&buf, &publicBlock); err != nil {
		return "", fmt.Errorf("encoding to public PEM: %w", err)
	}

	return buf.String(), nil
}
