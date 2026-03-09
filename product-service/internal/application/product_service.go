package application

import (
	"context"
	"github/productMCS/internal/domain"
	"github/productMCS/internal/ports/spi"

	"github.com/google/uuid"
)

type ProductService struct {
	repo spi.ProductRepository
}

func NewProductService(repo spi.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *domain.Product) error {

	if err := domain.ValidateProduct(product); err != nil {
		return err
	}

	return s.repo.Save(ctx, product)
}

func (s *ProductService) GetProduct(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ProductService) ListProducts(ctx context.Context) ([]*domain.Product, error) {
	return s.repo.FindAll(ctx)
}

func (s *ProductService) CheckStock(ctx context.Context, productID uuid.UUID, quantity int64) (bool, int64, error) {

	product, err := s.repo.FindByID(ctx, productID)
	if err != nil {
		return false, 0, err
	}

	if product.Stock >= quantity {
		return true, product.Stock, nil
	}

	return false, product.Stock, nil
}
