package middleware

import (
	"net/http"

	authentication "github.com/quinlanmorake/verisart-go/authentication"
	httpHelpers "github.com/quinlanmorake/verisart-go/helpers/http"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, authenticationCheck := authentication.GetUserFromTokenInHeaders(r); authenticationCheck.IsNotOk() {
			httpHelpers.WriteResponse(w, authenticationCheck)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
