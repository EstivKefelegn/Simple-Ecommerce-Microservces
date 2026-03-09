package domain

import (
	"errors"
	"strings"
)

func ValidateProduct(p *Product) error {

	if p == nil {
		return errors.New("product cannot be nil")
	}

	p.Name = strings.TrimSpace(p.Name)

	if p.Name == "" {
		return errors.New("product name required")
	}

	if p.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	if p.Stock < 0 {
		return errors.New("stock cannot be negative")
	}

	return nil
}