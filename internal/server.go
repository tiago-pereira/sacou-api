package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sacou-api/internal/users"

	"github.com/jackc/pgx/v5/pgxpool"
)

func StartServer() {
	// Router principal
	mux := http.NewServeMux()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/sacou_api"
	}

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("unable to connect to db: %v", err)
	}
	defer pool.Close()

	mux.Handle("/users", users.Init(pool))

	// Test endpoint raiz
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "API rodando!")
	})

	log.Println("Server running at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", mux))

}
