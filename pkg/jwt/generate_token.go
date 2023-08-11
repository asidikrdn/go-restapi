package jwtToken

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

// generate token
func GenerateToken(claims *jwt.MapClaims) (string, error) {
	// create new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the token with secret key
	webtoken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	// return signed token
	return webtoken, nil
}
