package spi

import (
	"context"
	"github/productMCS/internal/domain"

	"github.com/google/uuid"
)

type ProductRepository interface {
	Save(ctx context.Context, product *domain.Product) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Product, error)
	FindAll(ctx context.Context) ([]*domain.Product, error)
}
