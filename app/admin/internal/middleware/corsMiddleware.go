package middleware

import "net/http"

func CorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
	return handler
}
