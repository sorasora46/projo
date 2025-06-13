package routers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/projo/backend/config"
	"github.com/sorasora46/projo/backend/internal/adaptors/user"
	"github.com/sorasora46/projo/backend/internal/usecases"
	vldt "github.com/sorasora46/projo/backend/internal/validator"
)

func NewUserRoutes(api fiber.Router, database config.Database, envManager config.EnvManager, reqValidator vldt.ReqValidator) {
	db, err := database.GetDBInstance()
	if err != nil {
		log.Printf("[NewUserRoutes]: %v", err)
	}
	userRepository := user.NewUserRepository(db)
	userUsecases := usecases.NewUserUsercase(userRepository, envManager)
	userHandlers := user.NewUserHandler(userUsecases, reqValidator)

	api.Post("/", userHandlers.CreateUser)
	api.Post("/login", userHandlers.Login)
	api.Get("/:username", userHandlers.GetByUsername)
	api.Delete("/:username", userHandlers.DeleteByUsername)
}
