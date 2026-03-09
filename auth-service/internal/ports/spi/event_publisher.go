package spi

import (
	"context"
	"github/ecommerceMSAuth/internal/domain"
)

type EventPublisher interface {
	PublishUserCreated(context context.Context, event domain.UserCreatedEvent) error
}
