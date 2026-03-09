package application

import (
	"context"
	"errors"
	"github/orderService/internal/domain"
	"github/orderService/ports/spi"

	"github.com/google/uuid"
)

type OrderService struct {
	repo          spi.OrderRepository
	productClient spi.ProductClient
}

func NewOrderService(r spi.OrderRepository, pc spi.ProductClient) *OrderService {
	return &OrderService{
		repo:          r,
		productClient: pc,
	}
}
func (s *OrderService) CreateOrder(ctx context.Context, productID, userID uuid.UUID, quantity int) (uuid.UUID, error) {
    available, _, err := s.productClient.CheckStock(ctx, productID, quantity)
    if err != nil {
        return uuid.Nil, err
    }

    if !available {
        return uuid.Nil, errors.New("product out of stock")
    }

    order := &domain.Order{
        ID:        uuid.New(),
        ProductID: productID,
        UserID:    userID,
        Quantity:  quantity,
        Status:    "CREATED",
    }

    if err := s.repo.Save(ctx, order); err != nil {
        return uuid.Nil, err
    }

    return order.ID, nil
}