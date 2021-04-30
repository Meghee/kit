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
func HasValidJWT(notAuthorizedHandler http.HandlerFunc) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtToken, jwtPayload, err := retrieveJWTFromHeader(w, r, notAuthorizedHandler)
			if err != nil {
				notAuthorizedHandler(w, r)
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
// e.g JWTIndexIS("role", "admin", notAuthorizedHandler)
func JWTIndexIS(index string, value interface{}, notAuthorizedHandler http.HandlerFunc) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtToken, jwtPayload, err := retrieveJWTFromHeader(w, r, notAuthorizedHandler)
			if err != nil {
				notAuthorizedHandler(w, r)
				return
			}
			if jwtPayload[index] != value {
				notAuthorizedHandler(w, r)
				return
			}
			jwtCtx := context.WithValue(r.Context(), JWTCtxKey, jwtToken)
			jwtPayloadCtx := context.WithValue(jwtCtx, JWTPayloadCtxKey, jwtPayload)
			next.ServeHTTP(w, r.WithContext(jwtPayloadCtx))
		})
	}
}

// retrieveJWTFromHeader retrieves the jwt token from the http 'Authorization'
// header.
func retrieveJWTFromHeader(w http.ResponseWriter, r *http.Request, notAuthorizedHandler http.HandlerFunc) (string, map[string]interface{}, error) {
	authorization := r.Header.Get("Authorization")
	if len(authorization) < len("Bearer ")+1 {
		notAuthorizedHandler(w, r)
		return "", nil, errors.New("No jwt in authorization header")
	}
	jwtTokenString := authorization[len("Bearer "):]
	jwtClaims, err := hs256.Decode(jwtTokenString)
	if err != nil {
		notAuthorizedHandler(w, r)
		return "", nil, err
	}
	return jwtTokenString, jwtClaims, nil
}
