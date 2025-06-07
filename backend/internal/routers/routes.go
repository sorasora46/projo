package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/infras"
)

func InitRoutes(app *fiber.App, database infras.Database, envManager infras.EnvManager) {
	api := app.Group("/api")
	NewUserRoutes(api.Group("/user"), database, envManager)
}
