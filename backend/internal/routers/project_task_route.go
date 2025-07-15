package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/config"
	vldt "github.com/sorasora46/projo/backend/internal/validator"
)

func NewProjectTaskRoutes(api fiber.Router, database config.Database, envManager config.EnvManager, reqValidator vldt.ReqValidator) {
	// db, err := database.GetDBInstance()
	// if err != nil {
	// 	log.Printf("[NewProjectTaskRoutes]: %v", err)
	// }

	// projectRepo := project.NewProjectRepository(db)
	// projectUsecase := usecases.NewProjectUsecase(projectRepo)
	// projectHandlers := project.NewProjectHandler(projectUsecase, reqValidator)

	// TODO: group project task endpoint

	api.Post("/", func(c *fiber.Ctx) error {
		return nil
	})
	api.Get("/:taskId", func(c *fiber.Ctx) error {
		return nil
	})
	api.Patch("/:taskId", func(c *fiber.Ctx) error {
		return nil
	})
	api.Delete("/:taskId", func(c *fiber.Ctx) error {
		return nil
	})
}
