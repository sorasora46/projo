package routers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/internal/adaptors/user"
	"github.com/sorasora46/projo/backend/internal/infras"
	"github.com/sorasora46/projo/backend/internal/usecases"
)

func NewUserRoutes(api fiber.Router, database infras.Database) {
	db, err := database.GetDBInstance()
	if err != nil {
		log.Printf("[NewUserRoutes]: %v", err)
	}
	userRepository := user.NewUserRepository(db)
	userUsecases := usecases.NewUserUsercase(userRepository)
	userHandlers := user.NewUserHandler(userUsecases)

	api.Post("/", userHandlers.CreateUser)
	api.Get("/:username", userHandlers.GetByUsername)
	api.Delete("/:username", userHandlers.DeleteByUsername)
}
