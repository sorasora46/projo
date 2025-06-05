package routers

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	api := app.Group("/api")
	NewUserRoutes(api.Group("/user"))
}
