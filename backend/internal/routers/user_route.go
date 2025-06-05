package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/adaptors/user"
	"github.com/sorasora46/projo/backend/internal/infras"
	"github.com/sorasora46/projo/backend/internal/usecases"
)

func NewUserRoutes(api fiber.Router) {
	userRepository := user.NewUserRepository(infras.DB)
	userUsecases := usecases.NewUserUsercase(userRepository)
	userHandlers := user.NewUserHandler(userUsecases)

	api.Post("/", userHandlers.Create)
	api.Get("/:username", userHandlers.Get)
	api.Patch("/:username", userHandlers.Update)
	api.Delete("/:username", userHandlers.Delete)
}
