package jwtToken

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

// verify token
func VerifyToken(tokenString string) (*jwt.Token, error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)

	// parsing & validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	// if parsing is not successful, return error
	if err != nil {
		return nil, err
	}

	// if parsing is successful, return the token
	return token, nil
}

// decode token
func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	// verify token
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	// get claims form the token
	claims, isOK := token.Claims.(jwt.MapClaims)
	// verify the claims
	if !isOK || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	// return the claims
	return claims, nil
}
