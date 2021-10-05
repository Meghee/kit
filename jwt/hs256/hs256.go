package hs256

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// Encode encodes a jwt token using data gotten from payload.
func Encode(payload map[string]interface{}) (tokenString string, err error) {
	secretKey := getSecret()
	if len(secretKey) < 1 {
		return "", errors.New("No 'JWT_SECRET_KEY' value in environment variables")
	}
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
	secretKey := getSecret()
	if len(secretKey) < 1 {
		return nil, errors.New("No 'JWT_SECRET_KEY' value in environment variables")
	}
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

func getSecret() (secret []byte) {
	secret = []byte(os.Getenv("JWT_SECRET_KEY"))
	return secret
}
