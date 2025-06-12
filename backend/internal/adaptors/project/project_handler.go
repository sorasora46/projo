package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/dtos"
	"github.com/sorasora46/projo/backend/internal/usecases"
	"github.com/sorasora46/projo/backend/pkg/constants"
)

type ProjectHandler interface {
	CreateProject(c *fiber.Ctx) error
	GetByProjectId(c *fiber.Ctx) error
	GetAllProjects(c *fiber.Ctx) error
	DeleteByProjectId(c *fiber.Ctx) error
	UpdateProject(c *fiber.Ctx) error
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
		return c.Status(fiber.StatusBadRequest).JSON(dtos.CommonRes{
			Result: err.Error(),
		})
	}
	userId, ok := c.Locals(constants.UserIdContext).(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.CommonRes{
			Result: "error converting userId in Locals",
		})
	}
	if err := p.usecase.CreateProject(req, userId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.CommonRes{
			Result: err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(dtos.CommonRes{})
}

func (p *ProjectHandlerImpl) GetByProjectId(c *fiber.Ctx) error {
	projectId := c.Params(constants.ProjectIdParam)
	project, err := p.usecase.GetByProjectId(projectId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.CommonRes{
			Result: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(dtos.CommonRes{
		Result: project,
	})
}

func (p *ProjectHandlerImpl) GetAllProjects(c *fiber.Ctx) error {
	userId, ok := c.Locals(constants.UserIdContext).(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.CommonRes{
			Result: "error converting userId in Locals",
		})
	}
	projects, err := p.usecase.GetAllProjects(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.CommonRes{
			Result: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(dtos.CommonRes{
		Result: projects,
	})
}

func (p *ProjectHandlerImpl) DeleteByProjectId(c *fiber.Ctx) error {
	projectId := c.Params(constants.ProjectIdParam)
	err := p.usecase.DeleteByProjectId(projectId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.CommonRes{
			Result: err.Error(),
		})
	}
	return c.Status(fiber.StatusNoContent).JSON(dtos.CommonRes{})
}

func (p *ProjectHandlerImpl) UpdateProject(c *fiber.Ctx) error {
	projectId := c.Params(constants.ProjectIdParam)
	var req dtos.UpdateProjectReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.CommonRes{
			Result: err.Error(),
		})
	}
	if err := p.usecase.UpdateProject(req, projectId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.CommonRes{
			Result: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(dtos.CommonRes{})
}
