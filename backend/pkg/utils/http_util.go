package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/dtos"
	"github.com/sorasora46/projo/backend/pkg/constants"
)

func NewSuccessRes(c *fiber.Ctx, resp dtos.Response) error {
	body := dtos.SuccessRes{
		Success: constants.Success,
		Result:  resp.Result,
	}
	return c.Status(resp.Code).JSON(body)
}

func NewFailRes(c *fiber.Ctx, resp dtos.Response) error {
	body := dtos.FailRes{
		Success: constants.Fail,
		Message: resp.Error.Error(),
	}
	return c.Status(resp.Code).JSON(body)
}

func NewFailValidationRes(c *fiber.Ctx, code int, errs []dtos.ValidationError) error {
	body := dtos.FailValidationRes{
		Success: constants.Fail,
		Errors:  errs,
	}
	return c.Status(code).JSON(body)
}
