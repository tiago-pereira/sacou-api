.PHONY: migrate-up migrate-down sqlc-generate run dev-start

DB_URL = postgres://postgres:postgres@localhost:5432/sacou_api?sslmode=disable

migrate-up:
	migrate -path db/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path db/migrations -database "$(DB_URL)" down

dev-start:
	@echo "Starting docker containers..."
	docker compose up -d
	@echo "Waiting for Postgres to be ready..."
	@sleep 5
	$(MAKE) migrate-up
	$(MAKE) sqlc-generate
	$(MAKE) run

sqlc-generate:
	@echo "Generating Go code from SQL..."
	sqlc generate

run:
	@echo "Starting Go application..."
	go run cmd/main.go