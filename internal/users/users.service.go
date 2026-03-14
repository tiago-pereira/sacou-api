package users

import (
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, params CreateUserParams) (User, error)
	ListUsers(ctx context.Context) ([]User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) UserService {
	return &service{repo: repo}
}

func (s *service) CreateUser(ctx context.Context, params CreateUserParams) (User, error) {
	user := User{
		Name:  params.Name,
		Email: params.Email,
	}
	return s.repo.Create(ctx, user)
}

func (s *service) ListUsers(ctx context.Context) ([]User, error) {
	return s.repo.List(ctx)
}
