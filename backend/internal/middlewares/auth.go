package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sorasora46/projo/backend/internal/dtos"
	"github.com/sorasora46/projo/backend/internal/infras"
)

type AuthMiddleware interface {
	ValidateToken(c *fiber.Ctx) error
}

type AuthMiddlewareImpl struct {
	envManager infras.EnvManager
}

func NewAuthMiddleware(envManager infras.EnvManager) AuthMiddleware {
	return &AuthMiddlewareImpl{envManager: envManager}
}

func (a *AuthMiddlewareImpl) ValidateToken(c *fiber.Ctx) error {
	skipMap := map[string][]string{
		"POST": {"/api/user/", "/api/user/login"},
	}

	path := c.Path()
	method := c.Method()

	if skipPaths, ok := skipMap[method]; ok {
		for _, skipPath := range skipPaths {
			if path == skipPath {
				return c.Next()
			}
		}
	}

	accessToken := c.Cookies("accessToken")
	if len(accessToken) == 0 {
		return c.Status(401).JSON(dtos.CommonRes{})
	}

	_, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.envManager.GetJWTSignKey()), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS384.Alg()}))
	if err != nil {
		return c.Status(401).JSON(dtos.CommonRes{
			Result: err.Error(),
		})
	}
	return c.Next()
}
