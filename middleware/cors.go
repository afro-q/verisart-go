package middleware

import (
	"net/http"

	config "github.com/quinlanmorake/verisart-go/config"
)

const HTTP_OPTIONS = "OTIONS"

// Set headers for CORS
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setCors(w)

		if r.Method != HTTP_OPTIONS {
			next.ServeHTTP(w, r)
		}
	})
}

func setCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Methods", config.AppConfig.Cors.AllowMethods)
	w.Header().Set("Access-Control-Allow-Origin", config.AppConfig.Cors.AllowOrigin)
	w.Header().Set("Access-Control-Allow-Headers", config.AppConfig.Cors.AllowHeaders)
}
