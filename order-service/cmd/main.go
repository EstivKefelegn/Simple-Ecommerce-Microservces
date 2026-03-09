package main

import (
	"log"
	"net/http"

	grpcClient "github/orderService/adapters/grpc"
	https "github/orderService/adapters/http"
	"github/orderService/adapters/repository/postgres"
	"github/orderService/infrastructure/db"
	"github/orderService/internal/application"

	pb "github/orderService/product-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	dbConn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	orderRepo := postgres.NewOrderRepo(dbConn)

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}


	productClient := pb.NewProductServiceClient(conn)

	pc := grpcClient.NewProductClient(productClient)

	service := application.NewOrderService(orderRepo, pc)

	handler := https.NewHandler(service)
	router := https.Router(handler)

	log.Println("Order service running on :8082")
	log.Println("Grpc service running on :50051")
	err = http.ListenAndServe(":8082", router)
	if err != nil {
		log.Fatal(err)
	}
}