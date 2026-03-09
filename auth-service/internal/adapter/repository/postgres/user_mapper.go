package postgres

import "github/ecommerceMSAuth/internal/domain"

func ToDomain(u *User) *domain.User {
	if u == nil {
		return nil
	}

	return &domain.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}

func ToModel(u *domain.User) *User {
	if u == nil {
		return nil
	}

	return &User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}