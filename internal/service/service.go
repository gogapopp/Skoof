package service

import (
	"context"

	"github.com/gogapopp/Skoof/internal/model"
)

type Storager interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUser(ctx context.Context, emailOrUsername, password string) (model.User, error)
}

type Service struct {
	store Storager
}

func New(storage Storager) *Service {
	return &Service{
		store: storage,
	}
}

func (s *Service) CreateUser(ctx context.Context, user model.User) error {
	return s.store.CreateUser(ctx, user)
}
func (s *Service) GetUser(ctx context.Context, emailOrUsername, password string) (model.User, error) {
	return s.store.GetUser(ctx, emailOrUsername, password)
}
