package grpc

import (
	"context"
	"github/productMCS/internal/application"
	pb "github/productMCS/product-service/proto"
	"github.com/google/uuid"
)

type ProductGRPCServer struct {
	pb.UnimplementedProductServiceServer
	service *application.ProductService
}

func NewProductGRPCServer(service *application.ProductService) *ProductGRPCServer {
	return &ProductGRPCServer{
		service: service,
	}
}

func (s *ProductGRPCServer) CheckStock(ctx context.Context, req *pb.CheckStockRequest) (*pb.CheckStockResponse, error) {
	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, err
	}

	available, stock, err := s.service.CheckStock(ctx, productID, int64(req.Quantity))
	if err != nil {
		return nil, err
	}

	return &pb.CheckStockResponse{
		Available:    available,
		CurrentStock: int32(stock),
	}, nil
}