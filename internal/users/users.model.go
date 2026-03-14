package users

import (
	"time"

	"github.com/google/uuid"
)

// User represents the user entity
type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUserParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
