package middleware

import (
	"github/ecommerceMSCGateway/clients"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		valid := clients.ValidateToken(token)

		if !valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
