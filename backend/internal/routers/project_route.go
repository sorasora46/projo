package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/adaptors/project"
	"github.com/sorasora46/projo/backend/internal/infras"
)

func NewProjectRoutes(api fiber.Router, database infras.Database, envManager infras.EnvManager) {
	// db, err := database.GetDBInstance()
	// if err != nil {
	// 	log.Printf("[NewProjectRoutes]: %v", err)
	// }

	projectHandlers := project.NewProjectHandler()

	api.Post("/", projectHandlers.CreateProject)
	api.Get("/:projectId", projectHandlers.GetByProjectId)
}
