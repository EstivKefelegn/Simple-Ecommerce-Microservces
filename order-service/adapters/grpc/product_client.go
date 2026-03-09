package grpc

import (
	"context"
	"fmt"

	pb "github/orderService/product-service/proto"

	"github.com/google/uuid"
)

type ProductClient struct {
	client pb.ProductServiceClient
}

func NewProductClient(c pb.ProductServiceClient) *ProductClient {
	return &ProductClient{client: c}
}

func (p *ProductClient) CheckStock(ctx context.Context, productID uuid.UUID, quantity int) (bool, int, error) {

	req := &pb.CheckStockRequest{
		ProductId: productID.String(),
		Quantity:  int32(quantity),
	}
	resp, err := p.client.CheckStock(ctx, req)

	if err != nil {
		return false, 0, fmt.Errorf("stock check failed: %w", err)
	}

	return resp.Available, int(resp.CurrentStock), nil
}
