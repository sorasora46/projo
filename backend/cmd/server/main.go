package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/infras"
	"github.com/sorasora46/projo/backend/internal/routers"
)

func main() {
	app := fiber.New()

	// INIT ENV
	envManager := infras.NewEnvManager()
	envManager.InitEnv()

	// INIT DATABASE
	database := infras.NewDatabase()
	database.InitDB(envManager.DB_DSN)

	// INIT ROUTES
	routers.InitRoutes(app, database, *envManager)

	app.Listen(envManager.ADDR)
}
