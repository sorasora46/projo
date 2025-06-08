package project

import "github.com/gofiber/fiber/v2"

type ProjectHandler interface {
	CreateProject(c *fiber.Ctx) error
	GetByProjectId(c *fiber.Ctx) error
}

type ProjectHandlerImpl struct {
}

func NewProjectHandler() ProjectHandler {
	return &ProjectHandlerImpl{}
}

func (p *ProjectHandlerImpl) CreateProject(c *fiber.Ctx) error {
	return nil
}

func (p *ProjectHandlerImpl) GetByProjectId(c *fiber.Ctx) error {
	return nil
}
