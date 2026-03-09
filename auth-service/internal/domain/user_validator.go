package domain

import (
	"errors"
	"net/mail"
	"strings"

	"github/ecommerceMSAuth/pkg"
)

func ValidateUser(u *User) error {

	if u == nil {
		return errors.New("user cannot be nil")
	}

	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)

	if u.Username == "" {
		return errors.New("username cannot be empty")
	}

	if len(u.Username) < 3 {
		return errors.New("username must be at least 3 characters")
	}

	if u.Email == "" {
		return errors.New("email cannot be empty")
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return err
	}

	if len(u.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	u.ID = pkg.GenerateUUID()

	return nil
}

func ValidateEmail(email string) error {

	email = strings.TrimSpace(email)

	if email == "" {
		return errors.New("email cannot be empty")
	}

	_, err := mail.ParseAddress(email)
	return err
}
