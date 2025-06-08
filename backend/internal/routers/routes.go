package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/infras"
	"github.com/sorasora46/projo/backend/internal/middlewares"
)

func InitRoutes(app *fiber.App, database infras.Database, envManager infras.EnvManager) {
	authMiddleware := middlewares.NewAuthMiddleware(envManager)
	api := app.Group("/api", authMiddleware.ValidateToken)
	NewUserRoutes(api.Group("/user"), database, envManager)
	NewProjectRoutes(api.Group("/project"), database, envManager)
}
