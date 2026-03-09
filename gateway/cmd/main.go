// @title Ecommerce Microservices Gateway API
// @version 1.0
// @description API Gateway for Ecommerce Microservices (UserAuth, Product, Order). 
// @description It supports 3 communication ways: REST API to communicate the client frontend with our services via the gateway, gRPC to communicate service-to-service (e.g., check product stock when an order is created), RabbitMQ to store user-created events.
// @host localhost:8080
// @BasePath /
package main

import (
	"github/ecommerceMSCGateway/internal/router"
	"log"
	"net/http"
)

func main() {

	r := router.Routes()

	log.Println("Gateway running on :8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}