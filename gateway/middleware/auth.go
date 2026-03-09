package middleware

import (
	"context"
	"net/http"
	"strings"

	"github/ecommerceMSCGateway/clients"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token := parts[1]

		userID, valid := clients.ValidateTokenAndGetUserID(token)
		if !valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// Put user_id in context
		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
