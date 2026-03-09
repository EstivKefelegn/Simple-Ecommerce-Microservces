package api

import (
	"context"
	"github/productMCS/internal/domain"

	"github.com/google/uuid"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product *domain.Product) error
	GetProduct(ctx context.Context, id uuid.UUID) (*domain.Product, error)
	ListProducts(ctx context.Context) ([]*domain.Product, error)
}
