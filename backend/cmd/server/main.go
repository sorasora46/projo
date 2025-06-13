package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sorasora46/projo/backend/config"
	"github.com/sorasora46/projo/backend/internal/middlewares"
	"github.com/sorasora46/projo/backend/internal/routers"
	vldt "github.com/sorasora46/projo/backend/internal/validator"
)

func main() {
	// INIT ENV
	envManager := config.NewEnvManager()
	envManager.InitEnv()

	// INIT DATABASE
	database := config.NewDatabase()
	database.InitDB(envManager.GetDBDSN())

	// INIT VALIDATOR
	reqValidator := vldt.NewReqValidator(validator.New())

	// INIT ROUTES
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ErrorHandler:  middlewares.GlobalErrorHandler,
	})
	app.Use(recover.New())
	routers.InitRoutes(app, database, envManager, reqValidator)

	app.Listen(envManager.GetAddr())
}
