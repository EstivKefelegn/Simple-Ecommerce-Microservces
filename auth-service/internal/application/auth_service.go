package application

import (
	"context"
	"github/ecommerceMSAuth/internal/domain"
	"github/ecommerceMSAuth/internal/ports/spi"
	"github/ecommerceMSAuth/pkg"
	"time"
)

type AuthService struct {
	userRepo  spi.UserRepository
	publisher spi.EventPublisher
}

func NewAuthService(userRepo spi.UserRepository, publisher spi.EventPublisher) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		publisher: publisher,
	}
}

func (s *AuthService) Register(username, email, password string) (*domain.User, error) {
	user, _ := domain.NewUser(username, email, password)

	err := s.userRepo.Save(user)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	event := domain.UserCreatedEvent{ID: user.ID, Email: user.Email, CreatedAt: time.Now(), Version: 1}
	s.publisher.PublishUserCreated(ctx, event)

	return user, nil
}


func (s *AuthService) Login(email, password string) (string, error) {

	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if err := pkg.VerifyPassword(password, user.Password); err != nil {
		return "", err
	}

	token, err := pkg.SignToken(user.ID, user.Username, "user")
	if err != nil {
		return "", err
	}

	return token, nil
}
