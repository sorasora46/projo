package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/config"
	"github.com/sorasora46/projo/backend/internal/routers"
)

func main() {
	app := fiber.New()

	// INIT ENV
	envManager := config.NewEnvManager()
	envManager.InitEnv()

	// INIT DATABASE
	database := config.NewDatabase()
	database.InitDB(envManager.GetDBDSN())

	// INIT ROUTES
	routers.InitRoutes(app, database, envManager)

	app.Listen(envManager.GetAddr())
}
