package dtos

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/pkg/constants"
)

type FailRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   any    `json:"error"`
}

type SuccessRes struct {
	Success bool `json:"success"`
	Result  any  `json:"result"`
}

type Response struct {
	Code   int
	Result any
	Error  error
}

func NewSuccessRes(c *fiber.Ctx, resp Response) error {
	body := SuccessRes{
		Success: constants.Success,
		Result:  resp.Result,
	}
	return c.Status(resp.Code).JSON(body)
}

func NewFailRes(c *fiber.Ctx, resp Response) error {
	body := FailRes{
		Success: constants.Fail,
		Message: resp.Error.Error(),
		Error:   resp.Error,
	}
	return c.Status(resp.Code).JSON(body)
}
