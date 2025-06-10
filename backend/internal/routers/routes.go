package routers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/adaptors/user"
	"github.com/sorasora46/projo/backend/internal/infras"
	"github.com/sorasora46/projo/backend/internal/middlewares"
)

func InitRoutes(app *fiber.App, database infras.Database, envManager infras.EnvManager) {
	db, err := database.GetDBInstance()
	if err != nil {
		log.Fatal("[InitRoutes]: %v", err)
	}
	userRepo := user.NewUserRepository(db)
	authMiddleware := middlewares.NewAuthMiddleware(envManager, userRepo)
	api := app.Group("/api", authMiddleware.ValidateToken)
	NewUserRoutes(api.Group("/user"), database, envManager)
	NewProjectRoutes(api.Group("/project"), database, envManager)
}
