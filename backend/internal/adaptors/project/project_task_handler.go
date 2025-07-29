package project

import "github.com/gofiber/fiber/v2"

type ProjectTaskHandler interface {
	CreateTask(c *fiber.Ctx) error
	GetTaskById(c *fiber.Ctx) error
	UpdateTaskById(c *fiber.Ctx) error
	DeleteTaskById(c *fiber.Ctx) error
}

type ProjectTaskHandlerImpl struct {
}

func NewProjectTaskHandler() ProjectTaskHandler {
	return &ProjectTaskHandlerImpl{}
}

func (p *ProjectTaskHandlerImpl) CreateTask(c *fiber.Ctx) error {
	return c.SendString("CreateTask")
}

func (p *ProjectTaskHandlerImpl) GetTaskById(c *fiber.Ctx) error {
	return c.SendString("GetTaskByID")
}

func (p *ProjectTaskHandlerImpl) UpdateTaskById(c *fiber.Ctx) error {
	return c.SendString("UpdateTaskById")
}

func (p *ProjectTaskHandlerImpl) DeleteTaskById(c *fiber.Ctx) error {
	return c.SendString("DeleteTaskById")
}
