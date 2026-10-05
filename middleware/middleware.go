package middleware

import "net/http"

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Logging will be implemented later.

		next.ServeHTTP(w, r)
	})
}
