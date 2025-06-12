package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/dtos"
)

func GlobalErrorHandler(c *fiber.Ctx, err error) error {
	return dtos.NewFailRes(c, dtos.Response{
		Code:  fiber.StatusInternalServerError,
		Error: fiber.ErrInternalServerError,
	})
}
