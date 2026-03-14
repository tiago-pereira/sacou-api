package main

import (
	"sacou-api/internal"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	internal.StartServer()
}
