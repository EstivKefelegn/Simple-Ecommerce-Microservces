package http

import (
	"github.com/go-chi/chi/v5"
)

func Router(handler *AuthHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/auth/register", handler.Register)

	r.Post("/auth/login", handler.Login)

	r.Get("/validate", handler.Validate)
    return r
}

