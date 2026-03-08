package router

import (
	"github/ecommerceMSCGateway/handlers"
	"github/ecommerceMSCGateway/middleware"

	"github.com/go-chi/chi/v5"
)

func Routes() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logging)

	r.Post("/login", handlers.Login)

	r.Group(func(protected chi.Router) {

		protected.Use(middleware.AuthMiddleware)

		protected.Post("/orders", handlers.CreateOrder)

		protected.Get("/products", handlers.GetProducts)

	})

	return r
}
