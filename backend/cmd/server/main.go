package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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
			log.Printf("Warning: could not load .env.local: %v", err)
		}
	}

	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	ADDR := fmt.Sprintf("%s:%s", HOST, PORT)

	routers.InitRoutes(app)

	app.Listen(ADDR)
}
