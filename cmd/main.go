package main

import (
	"fiber-gorm/internal/infrastructure/database"
	"fiber-gorm/internal/infrastructure/http"
	"log"
)

func main() {
	// Initialize the database
	infrastructure.DatabaseInit()

	// Initialize and start the server
	server := http.NewServer()
	log.Fatal(server.Listen(":8000"))
}
