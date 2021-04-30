package chi

import (
	"net/http"

	"github.com/go-chi/chi"
)

// UrlParamsShouldNotBeEmpty returns a http middleware that checks if some url parameters
// are not empty.
//
// If one of the parameter is empty it returns a 404 error.
func UrlParamsShouldNotBeEmpty(notFoundHandler http.HandlerFunc, params ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, param := range params {
				parameter := chi.URLParam(r, param)
				if parameter == "" {
					notFoundHandler(w, r)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
