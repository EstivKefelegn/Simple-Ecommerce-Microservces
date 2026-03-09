package spi

import (
	"context"
	"github.com/google/uuid"
)

type ProductClient interface {
	CheckStock(ctx context.Context, productID uuid.UUID, quantity int) (bool, int, error)
}