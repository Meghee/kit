package hs256

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// Encode encodes a jwt token using data gotten from payload.
func Encode(payload map[string]interface{}) (tokenString string, err error) {
	claims := jwt.MapClaims{
		"iat": time.Now(),
	}
	for k, v := range payload {
		claims[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(secretKey)
	return
}

// Decode decodes a jwt token string.
//
// If the jwt token is invalid it returns an error.
func Decode(tokenString string) (claims map[string]interface{}, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.New("Invalid jwt token string")
		}
		return secretKey, nil
	})
	if err != nil {
		return
	}
	if token.Valid {
		claims = token.Claims.(jwt.MapClaims)
		return
	}
	err = errors.New("An unknowm error occured while decoding jwt")
	return
}
