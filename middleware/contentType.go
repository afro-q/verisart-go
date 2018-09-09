package middleware

/*
 Some clients don't set content type, and the server has trouble de-serializing
*/

import (
	"net/http"
)

func ContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setContentType(w)
		next.ServeHTTP(w, r)
	})
}

func setContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
