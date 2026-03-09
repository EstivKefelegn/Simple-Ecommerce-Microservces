package postgres

import (
	"github/productMCS/internal/domain"

	"github.com/google/uuid"
)

func ToDomain(m *ProductModel) *domain.Product {

	id, _ := uuid.Parse(m.ID)

	return &domain.Product{
		ID:          id,
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
		Stock:       m.Stock,
	}
}
