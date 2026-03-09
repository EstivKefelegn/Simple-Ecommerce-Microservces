package main

import (
	"log"
	"net"
	"net/http"

	"github/productMCS/internal/adapter/grpc"
	https "github/productMCS/internal/adapter/http"
	"github/productMCS/internal/adapter/repository/postgres"
	"github/productMCS/internal/application"
	"github/productMCS/internal/infrastructure/db"
	pb "github/productMCS/product-service/proto"

	"github.com/joho/godotenv"
	grpcServer "google.golang.org/grpc"
)

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	repo := postgres.NewProductRepo(database)
	service := application.NewProductService(repo)

	httpHandler := https.NewHandler(service)
	httpRouter := https.Router(httpHandler)

	go func() {
		log.Println("HTTP server running on :8083")
		if err := http.ListenAndServe(":8083", httpRouter); err != nil {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	grpcServer := grpcServer.NewServer()
	productGRPC := grpc.NewProductGRPCServer(service)

	pb.RegisterProductServiceServer(grpcServer, productGRPC)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("gRPC listen error: %v", err)
	}

	log.Println("gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
}
