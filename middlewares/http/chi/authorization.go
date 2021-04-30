package chi

import (
	"context"
	"errors"
	"net/http"

	"github.com/Meghee/kit/jwt/hs256"
)

type contextkey string

var (
	JWTCtxKey contextkey = "jwt"
)

// HasValidJWT is a middleware that checks if the request contains a valid jwt
// authorization bearer token.
func HasValidJWT(notAuthorizedHandler, next http.HandlerFunc) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtPayload, err := retrieveJWTFromHeader(w, r, notAuthorizedHandler)
			if err != nil {
				notAuthorizedHandler(w, r)
				return
			}
			ctx := context.WithValue(r.Context(), JWTCtxKey, jwtPayload)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// retrieveJWTFromHeader retrieves the jwt token from the http 'Authorization'
// header.
func retrieveJWTFromHeader(w http.ResponseWriter, r *http.Request, notAuthorizedHandler http.HandlerFunc) (map[string]interface{}, error) {
	authorization := r.Header.Get("Authorization")
	if len(authorization) < len("Bearer ")+1 {
		notAuthorizedHandler(w, r)
		return nil, errors.New("No jwt in authorization header")
	}
	jwtTokenString := authorization[len("Bearer "):]
	jwtClaims, err := hs256.Decode(jwtTokenString)
	if err != nil {
		notAuthorizedHandler(w, r)
		return nil, err
	}
	return jwtClaims, nil
}
