package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/config"
	"github.com/sorasora46/projo/backend/internal/adaptors/project"
	vldt "github.com/sorasora46/projo/backend/internal/validator"
)

func NewProjectTaskRoutes(api fiber.Router, database config.Database, envManager config.EnvManager, reqValidator vldt.ReqValidator) {
	// db, err := database.GetDBInstance()
	// if err != nil {
	// 	log.Printf("[NewProjectTaskRoutes]: %v", err)
	// }

	// projectRepo := project.NewProjectRepository(db)
	// projectUsecase := usecases.NewProjectUsecase(projectRepo)
	handlers := project.NewProjectTaskHandler()

	api.Post("/", handlers.CreateTask)
	api.Get("/:taskId", handlers.GetTaskById)
	api.Patch("/:taskId", handlers.UpdateTaskById)
	api.Delete("/:taskId", handlers.DeleteTaskById)
}
