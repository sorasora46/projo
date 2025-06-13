package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/dtos"
	"github.com/sorasora46/projo/backend/internal/dtos/req"
	"github.com/sorasora46/projo/backend/internal/usecases"
	vldt "github.com/sorasora46/projo/backend/internal/validator"
	"github.com/sorasora46/projo/backend/pkg/constants"
	"github.com/sorasora46/projo/backend/pkg/utils"
)

type ProjectHandler interface {
	CreateProject(c *fiber.Ctx) error
	GetByProjectId(c *fiber.Ctx) error
	GetAllProjects(c *fiber.Ctx) error
	DeleteByProjectId(c *fiber.Ctx) error
	UpdateProject(c *fiber.Ctx) error
}

type ProjectHandlerImpl struct {
	usecase      usecases.ProjectUsecase
	reqValidator vldt.ReqValidator
}

func NewProjectHandler(usecase usecases.ProjectUsecase, reqValidator vldt.ReqValidator) ProjectHandler {
	return &ProjectHandlerImpl{usecase: usecase, reqValidator: reqValidator}
}

func (p *ProjectHandlerImpl) CreateProject(c *fiber.Ctx) error {
	var req req.CreateProjectReq
	if err := c.BodyParser(&req); err != nil {
		return utils.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusBadRequest,
			Error: err,
		})
	}

	if errs := p.reqValidator.Validate(req); errs != nil {
		return utils.NewFailValidationRes(c, fiber.StatusBadRequest, errs)
	}

	userId, ok := c.Locals(constants.UserIdContext).(string)
	if !ok {
		return utils.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusInternalServerError,
			Error: fiber.NewError(fiber.StatusInternalServerError, constants.ErrConvertUserIdInContext),
		})
	}

	if err := p.usecase.CreateProject(req, userId); err != nil {
		return utils.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusInternalServerError,
			Error: err,
		})
	}

	return utils.NewSuccessRes(c, dtos.Response{
		Code:   fiber.StatusCreated,
		Result: nil,
	})
}

func (p *ProjectHandlerImpl) GetByProjectId(c *fiber.Ctx) error {
	projectId := c.Params(constants.ProjectIdParam)
	project, err := p.usecase.GetByProjectId(projectId)
	if err != nil {
		return utils.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusInternalServerError,
			Error: err,
		})
	}
	return utils.NewSuccessRes(c, dtos.Response{
		Code:   fiber.StatusOK,
		Result: project,
	})
}

func (p *ProjectHandlerImpl) GetAllProjects(c *fiber.Ctx) error {
	userId, ok := c.Locals(constants.UserIdContext).(string)
	if !ok {
		return utils.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusInternalServerError,
			Error: fiber.NewError(fiber.StatusInternalServerError, constants.ErrConvertUserIdInContext),
		})
	}

	projects, err := p.usecase.GetAllProjects(userId)
	if err != nil {
		return utils.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusInternalServerError,
			Error: err,
		})
	}

	return utils.NewSuccessRes(c, dtos.Response{
		Code:   fiber.StatusOK,
		Result: projects,
	})
}

func (p *ProjectHandlerImpl) DeleteByProjectId(c *fiber.Ctx) error {
	projectId := c.Params(constants.ProjectIdParam)
	err := p.usecase.DeleteByProjectId(projectId)
	if err != nil {
		return utils.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusInternalServerError,
			Error: err,
		})
	}

	return utils.NewSuccessRes(c, dtos.Response{
		Code:   fiber.StatusNoContent,
		Result: nil,
	})
}

func (p *ProjectHandlerImpl) UpdateProject(c *fiber.Ctx) error {
	projectId := c.Params(constants.ProjectIdParam)
	var req req.UpdateProjectReq
	if err := c.BodyParser(&req); err != nil {
		return utils.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusBadRequest,
			Error: err,
		})
	}
	if err := p.usecase.UpdateProject(req, projectId); err != nil {
		return utils.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusInternalServerError,
			Error: err,
		})
	}

	return utils.NewSuccessRes(c, dtos.Response{
		Code:   fiber.StatusOK,
		Result: nil,
	})
}
