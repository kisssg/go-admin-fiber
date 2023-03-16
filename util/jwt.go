package util

import (
	"encoding/json"
	"time"

	"github.com/cristalhq/jwt/v5"
)

const SecretKey = "secret"

func GenerateJwt(issuer string) (string, error) {
	// create a Signer (HMAC in this example)
	key := []byte(SecretKey)
	signer, err := jwt.NewSignerHS(jwt.HS256, key)

	// create claims (you can create your own, see: Example_BuildUserClaims)
	claims := &jwt.RegisteredClaims{
		Issuer:    issuer,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}

	// create a Builder
	builder := jwt.NewBuilder(signer)

	// and build a Token
	token, err := builder.Build(claims)
	return token.String(), err
}

func ParseJwt(cookieJwt string) (string, error) {
	// create a Verifier (HMAC in this example)
	key := []byte(SecretKey)
	verifier, err := jwt.NewVerifierHS(jwt.HS256, key)

	// parse and verify a token
	tokenBytes := []byte(cookieJwt)
	token, err := jwt.Parse(tokenBytes, verifier)

	if err != nil || token == nil || len(token.Bytes()) == 0 {
		return "", err
	}

	// get Registered claims
	var newClaims jwt.RegisteredClaims
	err = json.Unmarshal(token.Claims(), &newClaims)

	return newClaims.Issuer, err
}
