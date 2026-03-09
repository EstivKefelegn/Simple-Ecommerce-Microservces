package domain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Username string
	Email    string
	Password string
}

func NewUser(username, email, password string) (*User, error) {

	u := &User{
		Username: username,
		Email:    email,
		Password: password,
	}

	if err := ValidateUser(u); err != nil {
		return nil, err
	}

	u.ID = uuid.New()

	return u, nil
}