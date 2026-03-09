package postgres

import "github.com/google/uuid"

type OrderModel struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`   
	ProductID uuid.UUID `db:"product_id"`
	Quantity  int       `db:"quantity"`
	Status    string    `db:"status"`
}