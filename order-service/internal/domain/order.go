package domain

import "github.com/google/uuid"

type Order struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ProductID uuid.UUID
	Quantity  int
	Status    string
}
