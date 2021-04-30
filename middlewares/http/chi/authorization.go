package chi

// UrlParamAndJWTIndexIsExact checks if the value a url param is exactly the same as
// the value a the jwt payload index.
//
// Note: the jwt token is gotten from the authorization bearer header.
// func UrlParamAndJWTIndexIsExact(param, jwtIndex string, notAuthorizedHandler http.HandlerFunc) func(next http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			jwtToken, jwtPayload, err := httpMd.RetrieveJWTFromHTTPHeader(w, r, notAuthorizedHandler)
// 			if err
// 		})
// 	}
// }
