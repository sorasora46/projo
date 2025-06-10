package dtos

import "github.com/golang-jwt/jwt/v5"

type CustomClaim struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}
