package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/Meghee/kit/jwt/hs256"
)

type contextkey string

var (
	JWTCtxKey        contextkey = "jwt"
	JWTPayloadCtxKey contextkey = "jwt-payload"
)

// HasValidJWT is a middleware that checks if the request contains a valid jwt
// authorization bearer token.
func HasValidJWT(errHandler http.HandlerFunc) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtToken, jwtPayload, err := RetrieveJWTFromHTTPHeader(w, r, errHandler)
			if err != nil {
				return
			}
			jwtCtx := context.WithValue(r.Context(), JWTCtxKey, jwtToken)
			jwtPayloadCtx := context.WithValue(jwtCtx, JWTPayloadCtxKey, jwtPayload)
			next.ServeHTTP(w, r.WithContext(jwtPayloadCtx))
		})
	}
}

// JWTIndexIS is a middleware that checks if the decoded jwt payload in the authorization
// bearer header index matches a particular values.
//
// e.g JWTIndexIS("role", "admin", errHandler)
func JWTIndexIS(index string, value interface{}, errHandler http.HandlerFunc) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtPayload := r.Context().Value(JWTPayloadCtxKey).(map[string]interface{})
			if jwtPayload[index] != value {
				errHandler(w, r)
				return
			}
			next.ServeHTTP(w, r.WithContext(r.Context()))
		})
	}
}

// RetrieveJWTFromHTTPHeader retrieves the jwt token from the http 'Authorization'
// header.
func RetrieveJWTFromHTTPHeader(w http.ResponseWriter, r *http.Request, errHandler http.HandlerFunc) (string, map[string]interface{}, error) {
	authorization := r.Header.Get("Authorization")
	if len(authorization) < len("Bearer ")+1 {
		errHandler(w, r)
		return "", nil, errors.New("No jwt in authorization header")
	}
	jwtTokenString := authorization[len("Bearer "):]
	jwtClaims, err := hs256.Decode(jwtTokenString)
	if err != nil {
		errHandler(w, r)
		return "", nil, err
	}
	return jwtTokenString, jwtClaims, nil
}
