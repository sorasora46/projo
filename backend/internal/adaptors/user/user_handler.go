package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/dtos"
	"github.com/sorasora46/projo/backend/internal/dtos/req"
	"github.com/sorasora46/projo/backend/internal/usecases"
	vldt "github.com/sorasora46/projo/backend/internal/validator"
	"github.com/sorasora46/projo/backend/pkg/constants"
	"github.com/sorasora46/projo/backend/pkg/utils"
)

type UserHandler interface {
	CreateUser(c *fiber.Ctx) error
	GetByUsername(c *fiber.Ctx) error
	DeleteByUsername(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type UserHandlerImpl struct {
	usecase      usecases.UserUsecase
	reqValidator vldt.ReqValidator
}

func NewUserHandler(usecase usecases.UserUsecase, reqValidator vldt.ReqValidator) UserHandler {
	return &UserHandlerImpl{usecase: usecase, reqValidator: reqValidator}
}

func (u *UserHandlerImpl) CreateUser(c *fiber.Ctx) error {
	var req req.CreateUserReq
	if err := c.BodyParser(&req); err != nil {
		return utils.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusBadRequest,
			Error: err,
		})
	}

	if errs := u.reqValidator.Validate(req); errs != nil {
		return utils.NewFailValidationRes(c, fiber.StatusBadRequest, errs)
	}

	if err := u.usecase.CreateUser(req); err != nil {
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

func (u *UserHandlerImpl) GetByUsername(c *fiber.Ctx) error {
	username := c.Params(constants.UsernameParam)
	user, err := u.usecase.GetByUsername(username)
	if err != nil {
		return utils.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusInternalServerError,
			Error: err,
		})
	}

	return utils.NewSuccessRes(c, dtos.Response{
		Code:   fiber.StatusOK,
		Result: user,
	})
}

func (u *UserHandlerImpl) DeleteByUsername(c *fiber.Ctx) error {
	username := c.Params(constants.UsernameParam)
	err := u.usecase.DeleteByUsername(username)
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

func (u *UserHandlerImpl) Login(c *fiber.Ctx) error {
	var req req.LoginReq
	if err := c.BodyParser(&req); err != nil {
		return utils.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusBadRequest,
			Error: err,
		})
	}

	if errs := u.reqValidator.Validate(req); errs != nil {
		return utils.NewFailValidationRes(c, fiber.StatusBadRequest, errs)
	}

	jwt, err := u.usecase.Login(req.Username, req.Password)
	if err != nil {
		return utils.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusInternalServerError,
			Error: err,
		})
	}

	cookie := &fiber.Cookie{
		Name:     constants.AuthCookieName,
		Value:    *jwt,
		MaxAge:   constants.AuthCookieMaxAge,
		HTTPOnly: true,
		Secure:   true,
	}
	c.Cookie(cookie)

	return utils.NewSuccessRes(c, dtos.Response{
		Code:   fiber.StatusOK,
		Result: nil,
	})
}
