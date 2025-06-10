package routers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/adaptors/project"
	"github.com/sorasora46/projo/backend/internal/infras"
	"github.com/sorasora46/projo/backend/internal/usecases"
)

func NewProjectRoutes(api fiber.Router, database infras.Database, envManager infras.EnvManager) {
	db, err := database.GetDBInstance()
	if err != nil {
		log.Printf("[NewProjectRoutes]: %v", err)
	}

	projectRepo := project.NewProjectRepository(db)
	projectUsecase := usecases.NewProjectUsecase(projectRepo)
	projectHandlers := project.NewProjectHandler(projectUsecase)

	api.Post("/", projectHandlers.CreateProject)
	api.Get("/:projectId", projectHandlers.GetByProjectId)
	api.Get("/", projectHandlers.GetAllProjects)
	api.Delete("/:projectId", projectHandlers.DeleteByProjectId)
}
