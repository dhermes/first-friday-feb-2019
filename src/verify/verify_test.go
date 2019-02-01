package verify

import (
	"io/ioutil"
	"testing"
	"time"

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
func TestVerify(t *testing.T) {
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

	timestamp := time.Now()

	var isValid bool
	isValid, err = Verify(token, publicKeyBytes, timestamp)
	if err != nil {
		t.Errorf("Verify encountered an error: %v.", err)
		return
	}

	if !isValid {
		t.Error("JWT was invalid.")
		return
	}
}
