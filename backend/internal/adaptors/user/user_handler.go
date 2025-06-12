package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/dtos"
	"github.com/sorasora46/projo/backend/internal/usecases"
	"github.com/sorasora46/projo/backend/pkg/constants"
)

type UserHandler interface {
	CreateUser(c *fiber.Ctx) error
	GetByUsername(c *fiber.Ctx) error
	DeleteByUsername(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type UserHandlerImpl struct {
	usecase usecases.UserUsecase
}

func NewUserHandler(usecase usecases.UserUsecase) UserHandler {
	return &UserHandlerImpl{usecase: usecase}
}

func (u *UserHandlerImpl) CreateUser(c *fiber.Ctx) error {
	var req dtos.CreateUserReq
	if err := c.BodyParser(&req); err != nil {
		return dtos.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusBadRequest,
			Error: err,
		})
	}

	if err := u.usecase.CreateUser(req); err != nil {
		return dtos.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusInternalServerError,
			Error: err,
		})
	}

	return dtos.NewSuccessRes(c, dtos.Response{
		Code:   fiber.StatusCreated,
		Result: nil,
	})
}

func (u *UserHandlerImpl) GetByUsername(c *fiber.Ctx) error {
	username := c.Params(constants.UsernameParam)
	user, err := u.usecase.GetByUsername(username)
	if err != nil {
		return dtos.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusInternalServerError,
			Error: err,
		})
	}

	return dtos.NewSuccessRes(c, dtos.Response{
		Code:   fiber.StatusOK,
		Result: user,
	})
}

func (u *UserHandlerImpl) DeleteByUsername(c *fiber.Ctx) error {
	username := c.Params(constants.UsernameParam)
	err := u.usecase.DeleteByUsername(username)
	if err != nil {
		return dtos.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusInternalServerError,
			Error: err,
		})
	}

	return dtos.NewSuccessRes(c, dtos.Response{
		Code:   fiber.StatusNoContent,
		Result: nil,
	})
}

func (u *UserHandlerImpl) Login(c *fiber.Ctx) error {
	var req dtos.LoginReq
	if err := c.BodyParser(&req); err != nil {
		return dtos.NewFailRes(c, dtos.Response{
			Code:  fiber.StatusBadRequest,
			Error: err,
		})
	}

	jwt, err := u.usecase.Login(req.Username, req.Password)
	if err != nil {
		return dtos.NewFailRes(c, dtos.Response{
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

	return dtos.NewSuccessRes(c, dtos.Response{
		Code:   fiber.StatusOK,
		Result: nil,
	})
}
