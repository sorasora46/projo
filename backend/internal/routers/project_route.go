package routers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/config"
	"github.com/sorasora46/projo/backend/internal/adaptors/project"
	"github.com/sorasora46/projo/backend/internal/usecases"
	vldt "github.com/sorasora46/projo/backend/internal/validator"
)

func NewProjectRoutes(api fiber.Router, database config.Database, envManager config.EnvManager, reqValidator vldt.ReqValidator) {
	db, err := database.GetDBInstance()
	if err != nil {
		log.Printf("[NewProjectRoutes]: %v", err)
	}

	projectRepo := project.NewProjectRepository(db)
	projectUsecase := usecases.NewProjectUsecase(projectRepo)
	projectHandlers := project.NewProjectHandler(projectUsecase, reqValidator)

	api.Post("/", projectHandlers.CreateProject)
	api.Get("/:projectId", projectHandlers.GetByProjectId)
	api.Get("/", projectHandlers.GetAllProjects)
	api.Delete("/:projectId", projectHandlers.DeleteByProjectId)
	api.Patch("/:projectId", projectHandlers.UpdateProject)
}
