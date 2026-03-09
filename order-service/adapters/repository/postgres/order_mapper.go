package postgres

import (
	"github/orderService/internal/domain"
	"github/orderService/pkg"

	"github.com/google/uuid"
)

func toModel(order *domain.Order) *OrderModel {

	if order.ID == uuid.Nil {
		order.ID = pkg.GenerateUUID()
	}

	return &OrderModel{
		ID:        order.ID,
		UserID:    order.UserID,
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
		Status:    order.Status,
	}
}

func toDomain(model *OrderModel) *domain.Order {
	return &domain.Order{
		ID:        model.ID,
		UserID:    model.UserID,
		ProductID: model.ProductID,
		Quantity:  model.Quantity,
		Status:    model.Status,
	}
}
