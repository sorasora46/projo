package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/dtos"
	"github.com/sorasora46/projo/backend/internal/usecases"
)

type UserHandler interface {
	CreateUser(c *fiber.Ctx) error
	GetByUsername(c *fiber.Ctx) error
	DeleteByUsername(c *fiber.Ctx) error
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
		return c.Status(400).JSON(dtos.CommonRes{
			Result: err,
		})
	}
	if err := u.usecase.CreateUser(req); err != nil {
		return c.Status(500).JSON(dtos.CommonRes{
			Result: err,
		})
	}

	return c.Status(201).JSON(dtos.CommonRes{})
}

func (u *UserHandlerImpl) GetByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	user, err := u.usecase.GetByUsername(username)
	if err != nil {
		return c.Status(500).JSON(dtos.CommonRes{
			Result: err,
		})
	}

	return c.Status(200).JSON(dtos.CommonRes{
		Result: user,
	})
}

func (u *UserHandlerImpl) DeleteByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	err := u.usecase.DeleteByUsername(username)
	if err != nil {
		return c.Status(500).JSON(dtos.CommonRes{
			Result: err,
		})
	}

	return c.Status(204).JSON(dtos.CommonRes{})
}
