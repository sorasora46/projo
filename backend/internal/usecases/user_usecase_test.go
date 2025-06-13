package usecases_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/sorasora46/projo/backend/internal/dtos/req"
	"github.com/sorasora46/projo/backend/internal/entities"
	"github.com/sorasora46/projo/backend/internal/usecases"
	"github.com/sorasora46/projo/backend/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser(t *testing.T) {
	t.Run("successfully creates a user", func(t *testing.T) {
		// Arrange
		req := req.CreateUserReq{
			FirstName: "John",
			LastName:  "Doe",
			Username:  "john_doe",
			Email:     "john@example.com",
			Password:  "securePassword123",
		}
		mockRepo := new(mocks.MockUserRepository)
		mockRepo.On("Create", mock.AnythingOfType("*entities.User")).Return(nil)

		userUsecase := usecases.NewUserUsercase(mockRepo, nil)

		err := userUsecase.CreateUser(req)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("returns error if repo.Create fails", func(t *testing.T) {
		// Arrange
		req := req.CreateUserReq{
			FirstName: "John",
			LastName:  "Doe",
			Username:  "john_doe",
			Email:     "john@example.com",
			Password:  "securePassword123",
		}
		mockRepo := new(mocks.MockUserRepository)
		mockRepo.On("Create", mock.AnythingOfType("*entities.User")).Return(errors.New("db failure"))

		userUsecase := usecases.NewUserUsercase(mockRepo, nil)

		err := userUsecase.CreateUser(req)
		assert.Error(t, err)
		assert.Equal(t, "db failure", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestGetByUsername(t *testing.T) {
	t.Run("successfully get a user by username", func(t *testing.T) {
		id := uuid.NewString()
		username := "john_doe"
		expectedUser := &entities.User{
			Id:             id,
			FirstName:      "John",
			LastName:       "Doe",
			Username:       "john_doe",
			Email:          "john@example.com",
			HashedPassword: []byte{},
		}

		mockRepo := new(mocks.MockUserRepository)
		mockRepo.On("GetByUsername", username).Return(expectedUser, nil)

		userUsecase := usecases.NewUserUsercase(mockRepo, nil)

		dto, err := userUsecase.GetByUsername(username)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser.FirstName, dto.FirstName)
		assert.Equal(t, expectedUser.LastName, dto.LastName)
		assert.Equal(t, expectedUser.Username, dto.Username)
		assert.Equal(t, expectedUser.Email, dto.Email)

		mockRepo.AssertExpectations(t)
	})

	t.Run("returns error if repo.GetByUsername fails", func(t *testing.T) {
		username := "john_doe"
		mockRepo := new(mocks.MockUserRepository)
		mockRepo.On("GetByUsername", username).Return(nil, errors.New("repo failure"))

		userUsecase := usecases.NewUserUsercase(mockRepo, nil)

		dto, err := userUsecase.GetByUsername(username)
		assert.Nil(t, dto)
		assert.Error(t, err)
		assert.Equal(t, "repo failure", err.Error())

		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteByUsername(t *testing.T) {
	t.Run("successfully delete a user by username", func(t *testing.T) {
		username := "john_doe"

		mockRepo := new(mocks.MockUserRepository)
		mockRepo.On("DeleteByUsername", username).Return(nil)

		userUsecase := usecases.NewUserUsercase(mockRepo, nil)

		err := userUsecase.DeleteByUsername(username)
		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("returns error if repo.DeleteByUsername fails", func(t *testing.T) {
		username := "john_doe"

		mockRepo := new(mocks.MockUserRepository)
		mockRepo.On("DeleteByUsername", username).Return(errors.New("repo failure"))

		userUsecase := usecases.NewUserUsercase(mockRepo, nil)
		err := userUsecase.DeleteByUsername(username)
		assert.Error(t, err)
		assert.Equal(t, "repo failure", err.Error())

		mockRepo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	t.Run("login successfully", func(t *testing.T) {
		password := "very_strong_password"
		req := req.LoginReq{
			Username: "john_doe",
			Password: password,
		}
		userId := uuid.NewString()
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		assert.NoError(t, err)

		expectedUser := &entities.User{
			Id:             userId,
			FirstName:      "John",
			LastName:       "Doe",
			Username:       "john_doe",
			Email:          "john@example.com",
			HashedPassword: hashedPassword,
		}

		mockRepo := new(mocks.MockUserRepository)
		mockEnvMngr := new(mocks.MockEnvManager)
		mockRepo.On("GetLoginInfoByUsername", req.Username).Return(expectedUser, nil)

		// Assume that the Login method generates a non-empty JWT.
		userUsecase := usecases.NewUserUsercase(mockRepo, mockEnvMngr)
		jwt, err := userUsecase.Login(req.Username, req.Password)

		assert.NoError(t, err)
		assert.NotNil(t, jwt)
		assert.NotEmpty(t, *jwt)

		mockRepo.AssertExpectations(t)
	})

	t.Run("returns error when repo.GetLoginInfoByUsername fails", func(t *testing.T) {
		req := req.LoginReq{
			Username: "john_doe",
			Password: "password",
		}

		mockRepo := new(mocks.MockUserRepository)
		mockEnvMngr := new(mocks.MockEnvManager)
		mockRepo.On("GetLoginInfoByUsername", req.Username).Return(nil, errors.New("repo failure"))

		userUsecase := usecases.NewUserUsercase(mockRepo, mockEnvMngr)
		jwt, err := userUsecase.Login(req.Username, req.Password)

		assert.Error(t, err)
		assert.Nil(t, jwt)
		assert.Equal(t, "repo failure", err.Error())

		mockRepo.AssertExpectations(t)
	})

	t.Run("returns error when bcrypt.CompareHashAndPassword fails", func(t *testing.T) {
		password := "correct_password"
		incorrectPassword := "wrong_password"
		req := req.LoginReq{
			Username: "john_doe",
			Password: incorrectPassword,
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		assert.NoError(t, err)

		expectedUser := &entities.User{
			Id:             uuid.NewString(),
			FirstName:      "John",
			LastName:       "Doe",
			Username:       "john_doe",
			Email:          "john@example.com",
			HashedPassword: hashedPassword,
		}

		mockRepo := new(mocks.MockUserRepository)
		mockEnvMngr := new(mocks.MockEnvManager)
		mockRepo.On("GetLoginInfoByUsername", req.Username).Return(expectedUser, nil)

		userUsecase := usecases.NewUserUsercase(mockRepo, mockEnvMngr)
		jwt, err := userUsecase.Login(req.Username, req.Password)

		assert.Error(t, err)
		assert.Nil(t, jwt)

		mockRepo.AssertExpectations(t)
	})
}
