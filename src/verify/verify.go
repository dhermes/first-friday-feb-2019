package verify

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const resourceName = "urn:first-friday-feb-2019"

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
	parsed, err := Parse(token, publicKeyPEMBytes)
	if err != nil {
		return false, err
	}
	var iat, exp float64
	var iss, aud string
	var ok bool

	// TODO: Use `kid`.
	if iat, ok = parsed.Header["iat"].(float64); !ok {
		return false, errors.New("`iat` header is missing.")
	}
	if exp, ok = parsed.Header["exp"].(float64); !ok {
		return false, errors.New("`exp` header is missing.")
	}
	if iss, ok = parsed.Header["iss"].(string); !ok {
		return false, errors.New("`iss` header is missing.")
	}
	if aud, ok = parsed.Header["aud"].(string); !ok {
		return false, errors.New("`aud` header is missing.")
	}

	// Actually verify the values.
	if aud != resourceName {
		return false, errors.New("Invalid `aud` header.")
	}
	if iss != resourceName {
		return false, errors.New("Invalid `iss` header.")
	}

	lifetime := exp - iat
	if lifetime <= 0.0 || lifetime > 3600.0 {
		return false, errors.New("Invalid token lifetime.")
	}

	iatTimestamp := time.Unix(int64(iat), 0)
	if timestamp.Before(iatTimestamp) {
		// TODO: Allow 5 minute jitter.
		return false, errors.New("Token was issued (`iss`) in the future.")
	}
	expTimestamp := time.Unix(int64(exp), 0)
	if expTimestamp.Before(timestamp) {
		// TODO: Allow 5 minute jitter.
		return false, errors.New("Token has already expired (`exp`).")
	}

	return true, nil
}
