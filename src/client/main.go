package script

import (
	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/google/uuid"
	// request "github.com/imroc/req"
)

type Config struct {
	EXPIRATION_TIME		int
	CACHEKEYVALUE			string
	GIMMEURL					string
	VERIFYURL					string
	ISSUER						string
}

func createJWT(config *Config, privateKey) {
	jwtToken := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		jwt.StandardClaims{
			Audience:			config.ISSUER,
			ExpiresAt: 		config.EXPIRATION_TIME,
    	Id:						uuid.NewV4(),
    	Issuer:				config.ISSUER,
    	Subject:			uuid.NewV4(),
		},
	)
	return jwtToken
}

func privateKey() {

}

func verifyIdentity(token) {
	
}

func script() {
	config := Config{
		EXPIRATION_TIME:	300,
		CACHEKEYVALUE:		"privateKey",
		GIMMEURL:					"",
		VERIFYURL:				"",
		ISSUER:						"urn:first-friday-feb-2019",
	}
	privateKey := getPkey()
  token := createJWT(privateKey)
  verifyIdentity(token)
}