package verify

import (
	"io/ioutil"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
)

func TestParse(t *testing.T) {
	tokenBytes, err := ioutil.ReadFile("test_data/example_jwt.txt")
	if err != nil {
		t.Errorf("Failed to load example JWT: %v.", err)
		return
	}
	token := string(tokenBytes) // Should check ASCII encoding.

	publicKeyBytes, err := ioutil.ReadFile("test_data/id_rsa.pub.pem")
	if err != nil {
		t.Errorf("Failed to load example public key: %v.", err)
		return
	}

	var parsed *jwt.Token
	parsed, err = Parse(token, publicKeyBytes)
	if err != nil {
		t.Errorf("Failed to parse example JWT: %v.", err)
		return
	}
	if parsed == nil {
		t.Error("Parsed token is `nil`.")
		return
	}

	if !parsed.Valid {
		t.Errorf("JWT was invalid.")
		return
	}
}
