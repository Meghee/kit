package chi

import (
	"net/http"

	httpMd "github.com/Meghee/kit/middlewares/http"
	"github.com/go-chi/chi"
)

// UrlParamAndJWTIndexIsExact checks if the value a url param is exactly the same as
// the value a the jwt payload index.
//
// Note: the jwt token is gotten from the authorization bearer header.
func UrlParamAndJWTIndexIsExact(param, jwtIndex string, notAuthorizedHandler http.HandlerFunc) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtPayload := r.Context().Value(httpMd.JWTPayloadCtxKey).(map[string]interface{})
			paramValue := chi.URLParam(r, param)
			if jwtPayload[jwtIndex] != paramValue {
				notAuthorizedHandler(w, r)
				return
			}
			next.ServeHTTP(w, r.WithContext(r.Context()))
		})
	}
}
