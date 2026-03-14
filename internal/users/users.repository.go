package users

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetByID(ctx context.Context, id uuid.UUID) (User, error)
	Create(ctx context.Context, user User) (User, error)
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context) ([]User, error)
}
