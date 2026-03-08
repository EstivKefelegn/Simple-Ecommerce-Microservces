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
