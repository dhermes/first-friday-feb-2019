package verify

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func Parse(token string, publicKeyPEMBytes []byte) (*jwt.Token, error) {
	derBlock, _ := pem.Decode(publicKeyPEMBytes)
	if derBlock == nil {
		return nil, errors.New("Failed to decode PEM block containing public key.")
	}

	publicKey, err := x509.ParsePKCS1PublicKey(derBlock.Bytes)
	if err != nil {
		return nil, err
	}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	}

	parsed, err := jwt.Parse(token, keyFunc)
	if err != nil {
		return nil, err
	}
	return parsed, nil
}

func Verify(token string, publicKeyPEMBytes []byte, timestamp time.Time) (bool, error) {
	_, err := Parse(token, publicKeyPEMBytes)
	if err != nil {
		return false, err
	}
	return true, nil
}
