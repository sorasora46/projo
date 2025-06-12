package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sorasora46/projo/backend/config"
	"github.com/sorasora46/projo/backend/internal/middlewares"
	"github.com/sorasora46/projo/backend/internal/routers"
)

func main() {
	// INIT ENV
	envManager := config.NewEnvManager()
	envManager.InitEnv()

	// INIT DATABASE
	database := config.NewDatabase()
	database.InitDB(envManager.GetDBDSN())

	// INIT ROUTES
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ErrorHandler:  middlewares.GlobalErrorHandler,
	})
	app.Use(recover.New())
	routers.InitRoutes(app, database, envManager)

	app.Listen(envManager.GetAddr())
}
