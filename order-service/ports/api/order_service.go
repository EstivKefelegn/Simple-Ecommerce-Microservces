package api

import (
	"context"
	"github.com/google/uuid"
)

type OrderService interface {
	CreateOrder(ctx context.Context, productID uuid.UUID, quantity int) error
}