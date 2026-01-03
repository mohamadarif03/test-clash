package main

import (
	"log"
	"test-clash-be/internal/router"
	"test-clash-be/pkg/database"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	database.Connect()
	database.Migrate()
	database.SeedAdmin()

	r := router.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
