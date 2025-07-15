package routers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/config"
	"github.com/sorasora46/projo/backend/internal/adaptors/user"
	"github.com/sorasora46/projo/backend/internal/middlewares"
	vldt "github.com/sorasora46/projo/backend/internal/validator"
)

func InitRoutes(app *fiber.App, database config.Database, envManager config.EnvManager, reqValidator vldt.ReqValidator) {
	db, err := database.GetDBInstance()
	if err != nil {
		log.Fatalf("[InitRoutes]: %v", err)
	}
	userRepo := user.NewUserRepository(db)
	authMiddleware := middlewares.NewAuthMiddleware(envManager, userRepo)
	api := app.Group("/api", authMiddleware.ValidateToken)
	NewUserRoutes(api.Group("/user"), database, envManager, reqValidator)
	NewProjectRoutes(api.Group("/project"), database, envManager, reqValidator)
	NewProjectTaskRoutes(api.Group("/:projectId/task"), database, envManager, reqValidator)
}
