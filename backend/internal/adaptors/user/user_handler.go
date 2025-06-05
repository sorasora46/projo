package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/usecases"
)

type UserHandler interface {
	Create(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type UserHandlerImpl struct {
	usecase usecases.UserUsecase
}

func NewUserHandler(usecase usecases.UserUsecase) UserHandler {
	return &UserHandlerImpl{usecase: usecase}
}

func (u *UserHandlerImpl) Create(c *fiber.Ctx) error { return nil }
func (u *UserHandlerImpl) Get(c *fiber.Ctx) error    { return nil }
func (u *UserHandlerImpl) Update(c *fiber.Ctx) error { return nil }
func (u *UserHandlerImpl) Delete(c *fiber.Ctx) error { return nil }
