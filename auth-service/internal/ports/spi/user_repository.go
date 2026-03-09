package spi

import "github/ecommerceMSAuth/internal/domain"

type UserRepository interface {
	Save(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}

