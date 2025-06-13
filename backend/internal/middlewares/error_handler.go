package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/dtos"
	"github.com/sorasora46/projo/backend/pkg/utils"
)

func GlobalErrorHandler(c *fiber.Ctx, err error) error {
	return utils.NewFailRes(c, dtos.Response{
		Code:  fiber.StatusInternalServerError,
		Error: fiber.NewError(fiber.StatusInternalServerError, err.Error()),
	})
}
