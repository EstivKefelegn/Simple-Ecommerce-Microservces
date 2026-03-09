package http

import (
	"github.com/go-chi/chi/v5"
)

func Router(handler *Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/products", handler.CreateProduct)
	r.Get("/products", handler.ListProducts)
	r.Get("/products/{id}", handler.GetProduct)

	return r
}