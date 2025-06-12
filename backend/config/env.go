package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvManager interface {
	InitEnv()
	GetAddr() string
	GetDBDSN() string
	GetJWTSignKey() string
}

type EnvManagerImpl struct {
	ENV          string
	HOST         string
	PORT         string
	ADDR         string
	DB_HOST      string
	DB_PORT      string
	DB_USERNAME  string
	DB_PASSWORD  string
	DB_NAME      string
	DB_DSN       string
	JWT_SIGN_KEY string
}

func NewEnvManager() EnvManager {
	return &EnvManagerImpl{}
}

func (e *EnvManagerImpl) InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	e.ENV = os.Getenv("ENV")

	if e.ENV == "DEV" {
		err = godotenv.Overload(".env.local")
		if err != nil {
			log.Printf("could not load .env.local: %v", err)
		}
	}

	e.readServerEnv()
	e.readDBEnv()
	e.readCredentialEnv()
}

func (e *EnvManagerImpl) readServerEnv() {
	e.HOST = os.Getenv("HOST")
	e.PORT = os.Getenv("PORT")
	e.ADDR = fmt.Sprintf("%s:%s", e.HOST, e.PORT)
}

func (e *EnvManagerImpl) readDBEnv() {
	e.DB_HOST = os.Getenv("DB_HOST")
	e.DB_PORT = os.Getenv("DB_PORT")
	e.DB_USERNAME = os.Getenv("DB_USERNAME")
	e.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	e.DB_NAME = os.Getenv("DB_NAME")
	e.DB_DSN = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		e.DB_HOST, e.DB_PORT, e.DB_USERNAME, e.DB_PASSWORD, e.DB_NAME,
	)
}

func (e *EnvManagerImpl) readCredentialEnv() {
	e.JWT_SIGN_KEY = os.Getenv("JWT_SIGN_KEY")
}

func (e *EnvManagerImpl) GetAddr() string {
	return e.ADDR
}

func (e *EnvManagerImpl) GetDBDSN() string {
	return e.DB_DSN
}

func (e *EnvManagerImpl) GetJWTSignKey() string {
	return e.JWT_SIGN_KEY
}
