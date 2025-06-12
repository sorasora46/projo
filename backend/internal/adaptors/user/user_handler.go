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
		return c.Status(fiber.StatusBadRequest).JSON(dtos.CommonRes{
			Result: err.Error(),
		})
	}
	if err := u.usecase.CreateUser(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.CommonRes{
			Result: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dtos.CommonRes{})
}

func (u *UserHandlerImpl) GetByUsername(c *fiber.Ctx) error {
	username := c.Params(constants.UsernameParam)
	user, err := u.usecase.GetByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.CommonRes{
			Result: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dtos.CommonRes{
		Result: user,
	})
}

func (u *UserHandlerImpl) DeleteByUsername(c *fiber.Ctx) error {
	username := c.Params(constants.UsernameParam)
	err := u.usecase.DeleteByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.CommonRes{
			Result: err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(dtos.CommonRes{})
}

func (u *UserHandlerImpl) Login(c *fiber.Ctx) error {
	var req dtos.LoginReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.CommonRes{
			Result: err.Error(),
		})
	}
	jwt, err := u.usecase.Login(req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.CommonRes{
			Result: err.Error(),
		})
	}

	cookie := &fiber.Cookie{
		Name:     constants.JwtCookieName,
		Value:    *jwt,
		MaxAge:   constants.JwtMaxAge,
		HTTPOnly: true,
		Secure:   true,
	}
	c.Cookie(cookie)

	return c.Status(fiber.StatusOK).JSON(dtos.CommonRes{})
}
