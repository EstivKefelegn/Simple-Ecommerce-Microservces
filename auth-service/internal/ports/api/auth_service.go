package api

import "github/ecommerceMSAuth/internal/domain"

type AuthUseCase interface {
	Register(username, email, password string) (*domain.User, error)
	Login(email string, password string) (string, error)

}

