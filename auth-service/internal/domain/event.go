package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserCreatedEvent struct {
	ID        uuid.UUID `json:"id"`
	Email     string `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	Version   int 	 `json:"version"`
}

