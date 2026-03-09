package main

import (
	"fmt"
	https "github/ecommerceMSAuth/internal/adapter/http"
	"github/ecommerceMSAuth/internal/adapter/messaging"
	"github/ecommerceMSAuth/internal/adapter/repository/postgres"
	"github/ecommerceMSAuth/internal/application"
	"github/ecommerceMSAuth/internal/infrastructure/db"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	db, _ := db.ConnectDB()
	PORT := os.Getenv("API_PORT")

	userRepo := postgres.NewPostgresUserRepo(db)
	rabbitConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	publisher, err := messaging.NewRabbitMQPublisher(rabbitConn)
	if err != nil {
		log.Fatal(err)
	}

	defer publisher.Close()

	authService := application.NewAuthService(userRepo, publisher)
	handler := https.NewAuthHandler(authService)

	router := https.Router(handler)
	fmt.Printf("Server running on port %v", PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", PORT), router)
}
