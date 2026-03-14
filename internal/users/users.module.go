package users

import (
	"net/http"
	"sacou-api/internal/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Init(pool *pgxpool.Pool) http.Handler {
	queries := db.New(pool)

	userRepo := NewSQLCRepository(queries)
	userService := NewService(userRepo)

	router := Router(userService)

	return router
}
