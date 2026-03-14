package users

import (
	"context"
	"sacou-api/internal/db"

	"github.com/google/uuid"
)

type sqlcRepository struct {
	queries *db.Queries
}

func NewSQLCRepository(queries *db.Queries) Repository {
	return &sqlcRepository{queries: queries}
}

func (r *sqlcRepository) Create(ctx context.Context, u User) (User, error) {
	created, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Name:  u.Name,
		Email: u.Email,
	})
	if err != nil {
		return User{}, err
	}
	return mapToDomain(created), nil
}

func (r *sqlcRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.queries.DeleteUserByID(ctx, id)
}

func (r *sqlcRepository) GetByID(ctx context.Context, id uuid.UUID) (User, error) {
	u, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return User{}, err
	}
	return mapToDomain(u), nil
}

func (r *sqlcRepository) List(ctx context.Context) ([]User, error) {
	users, err := r.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	var domainUsers []User
	for _, u := range users {
		domainUsers = append(domainUsers, mapToDomain(u))
	}
	return domainUsers, nil
}

// mapping sqlc -> domain
func mapToDomain(u db.User) User {
	return User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.Time,
	}
}
