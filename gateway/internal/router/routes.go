package router

import (
	"github/ecommerceMSCGateway/handlers"
	"github/ecommerceMSCGateway/middleware"

	"github.com/go-chi/chi/v5"
)

func Routes() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logging)

	r.Route("/auth", func(auth chi.Router) {
		auth.Post("/login", handlers.Login)
		auth.Post("/register", handlers.Register)
	})

	r.Get("/products", handlers.GetProducts)
	r.Get("/product/{id}", handlers.GetProduct)
	r.Group(func(protected chi.Router) {
		
		protected.Use(middleware.AuthMiddleware)

		protected.Post("/orders", handlers.CreateOrder)

		protected.Post("/products", handlers.CreateProduct)

	})

	return r
}
