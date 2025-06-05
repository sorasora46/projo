package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sorasora46/projo/backend/internal/infras"
	"github.com/sorasora46/projo/backend/internal/routers"
)

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ENV := os.Getenv("ENV")

	if ENV == "DEV" {
		err = godotenv.Overload(".env.local")
		if err != nil {
			log.Printf("could not load .env.local: %v", err)
		}
	}

	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	ADDR := fmt.Sprintf("%s:%s", HOST, PORT)

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD, DB_NAME,
	)
	infras.InitDB(DSN)

	routers.InitRoutes(app)

	app.Listen(ADDR)
}
