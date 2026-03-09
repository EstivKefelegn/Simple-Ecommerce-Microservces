package spi

import (
	"context"
	"github/orderService/internal/domain"
)

type OrderRepository interface {
	Save(ctx context.Context, order *domain.Order) error
}
