package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/dtos"
	"github.com/sorasora46/projo/backend/internal/usecases"
)

type ProjectHandler interface {
	CreateProject(c *fiber.Ctx) error
	GetByProjectId(c *fiber.Ctx) error
	GetAllProjects(c *fiber.Ctx) error
}

type ProjectHandlerImpl struct {
	usecase usecases.ProjectUsecase
}

func NewProjectHandler(usecase usecases.ProjectUsecase) ProjectHandler {
	return &ProjectHandlerImpl{usecase: usecase}
}

func (p *ProjectHandlerImpl) CreateProject(c *fiber.Ctx) error {
	var req dtos.CreateProjectReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(401).JSON(dtos.CommonRes{
			Result: err.Error(),
		})
	}
	// TODO: get userId from JWT
	if err := p.usecase.CreateProject(req, "userId"); err != nil {
		return c.Status(500).JSON(dtos.CommonRes{
			Result: err.Error(),
		})
	}
	return c.Status(201).JSON(dtos.CommonRes{})
}

func (p *ProjectHandlerImpl) GetByProjectId(c *fiber.Ctx) error {
	projectId := c.Params("projectId")
	project, err := p.usecase.GetByProjectId(projectId)
	if err != nil {
		return c.Status(500).JSON(dtos.CommonRes{
			Result: err.Error(),
		})
	}
	return c.Status(200).JSON(dtos.CommonRes{
		Result: project,
	})
}

func (p *ProjectHandlerImpl) GetAllProjects(c *fiber.Ctx) error {
	return nil
}
