package postgres

import (
	"context"

	"github/orderService/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Save(ctx context.Context, order *domain.Order) error {

	model := toModel(order)

	query := `
	INSERT INTO orders(id, user_id, product_id, quantity,status)
	 VALUES($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		model.ID,
		model.UserID,
		model.ProductID,
		model.Quantity,
		model.Status,
	)

	return err
}
