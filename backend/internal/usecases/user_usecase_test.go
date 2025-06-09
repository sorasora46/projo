package usecases_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/sorasora46/projo/backend/internal/dtos"
	"github.com/sorasora46/projo/backend/internal/entities"
	"github.com/sorasora46/projo/backend/internal/usecases"
	"github.com/sorasora46/projo/backend/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser(t *testing.T) {
	t.Run("successfully creates a user", func(t *testing.T) {
		mockRepo := &mocks.MockUserRepository{
			CreateFunc: func(user *entities.User) error {
				assert.Equal(t, "john_doe", user.Username)
				assert.Equal(t, "John", user.FirstName)
				assert.Equal(t, "Doe", user.LastName)
				assert.Equal(t, "john@example.com", user.Email)
				assert.NotEmpty(t, user.HashedPassword)
				assert.NotEqual(t, "securePassword123", user.HashedPassword)
				return nil
			},
		}

		userUsecase := usecases.NewUserUsercase(mockRepo, nil)

		req := dtos.CreateUserReq{
			FirstName: "John",
			LastName:  "Doe",
			Username:  "john_doe",
			Email:     "john@example.com",
			Password:  "securePassword123",
		}

		err := userUsecase.CreateUser(req)
		assert.NoError(t, err)
	})

	t.Run("returns error if repo.Create fails", func(t *testing.T) {
		mockRepo := &mocks.MockUserRepository{
			CreateFunc: func(user *entities.User) error {
				return errors.New("repo failure")
			},
		}

		userUsecase := usecases.NewUserUsercase(mockRepo, nil)

		req := dtos.CreateUserReq{
			FirstName: "Jane",
			LastName:  "Smith",
			Username:  "jane_smith",
			Email:     "jane@example.com",
			Password:  "anotherPassword456",
		}

		err := userUsecase.CreateUser(req)
		assert.Error(t, err)
		assert.Equal(t, "repo failure", err.Error())
	})
}

func TestGetByUsername(t *testing.T) {
	t.Run("successfully get a user by username", func(t *testing.T) {
		id := uuid.NewString()
		username := "john_doe"
		mockRepo := &mocks.MockUserRepository{
			GetByUsernameFunc: func(username string) (*entities.User, error) {
				return &entities.User{
					Id:             id,
					FirstName:      "John",
					LastName:       "Doe",
					Username:       "john_doe",
					Email:          "john@example.com",
					HashedPassword: []byte{},
				}, nil
			},
		}

		userUsecase := usecases.NewUserUsercase(mockRepo, nil)

		dto, err := userUsecase.GetByUsername(username)
		assert.Equal(t, "John", dto.FirstName)
		assert.Equal(t, "Doe", dto.LastName)
		assert.Equal(t, username, dto.Username)
		assert.Equal(t, "john@example.com", dto.Email)
		assert.NoError(t, err)
	})

	t.Run("returns error if repo.GetByUsername fails", func(t *testing.T) {
		username := "john_doe"
		mockRepo := &mocks.MockUserRepository{
			GetByUsernameFunc: func(username string) (*entities.User, error) {

				return nil, errors.New("repo failure")
			},
		}

		userUsecase := usecases.NewUserUsercase(mockRepo, nil)

		dto, err := userUsecase.GetByUsername(username)
		assert.Nil(t, dto)
		assert.Error(t, err)
		assert.Equal(t, "repo failure", err.Error())
	})
}

func TestDeleteByUsername(t *testing.T) {
	t.Run("successfully delete a user by username", func(t *testing.T) {
		username := "john_doe"
		mockRepo := &mocks.MockUserRepository{
			DeleteByUsernameFunc: func(username string) error {
				return nil
			},
		}

		userUsecase := usecases.NewUserUsercase(mockRepo, nil)

		err := userUsecase.DeleteByUsername(username)
		assert.Nil(t, err)
	})

	t.Run("returns error if repo.DeleteByUsername fails", func(t *testing.T) {
		username := "john_doe"
		mockRepo := &mocks.MockUserRepository{
			DeleteByUsernameFunc: func(username string) error {
				return errors.New("repo failure")
			},
		}

		userUsecase := usecases.NewUserUsercase(mockRepo, nil)
		err := userUsecase.DeleteByUsername(username)
		assert.Error(t, err)
		assert.Equal(t, "repo failure", err.Error())
	})
}

func TestLogin(t *testing.T) {
	t.Run("login successfully", func(t *testing.T) {
		password := "very_strong_password"
		req := dtos.LoginReq{
			Username: "john_doe",
			Password: password,
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		assert.NoError(t, err)

		mockEnvMngr := &mocks.MockEnvManager{}
		mockRepo := &mocks.MockUserRepository{
			GetHashedPasswordByUsernameFunc: func(username string) ([]byte, error) {
				return hashedPassword, nil
			},
		}

		userUsecase := usecases.NewUserUsercase(mockRepo, mockEnvMngr)

		jwt, err := userUsecase.Login(req.Username, req.Password)
		assert.NoError(t, err)
		assert.NotNil(t, jwt)
		assert.NotEmpty(t, *jwt)
	})

	t.Run("returns error when repo.GetHashedPasswordByUsername fails", func(t *testing.T) {
		password := "very_strong_password"
		req := dtos.LoginReq{
			Username: "john_doe",
			Password: password,
		}

		mockEnvMngr := &mocks.MockEnvManager{}
		mockRepo := &mocks.MockUserRepository{
			GetHashedPasswordByUsernameFunc: func(username string) ([]byte, error) {
				return nil, errors.New("repo failure")
			},
		}

		userUsecase := usecases.NewUserUsercase(mockRepo, mockEnvMngr)

		jwt, err := userUsecase.Login(req.Username, req.Password)
		assert.Error(t, err)
		assert.Equal(t, "repo failure", err.Error())
		assert.Nil(t, jwt)
	})

	t.Run("return errors when bcrypt.CompareHashAndPassword fails", func(t *testing.T) {
		incorrectPassword := "incorrect_password"
		password := "very_strong_password"
		req := dtos.LoginReq{
			Username: "john_doe",
			Password: incorrectPassword,
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		assert.NoError(t, err)

		mockEnvMngr := &mocks.MockEnvManager{}
		mockRepo := &mocks.MockUserRepository{
			GetHashedPasswordByUsernameFunc: func(username string) ([]byte, error) {
				return hashedPassword, nil
			},
		}

		userUsecase := usecases.NewUserUsercase(mockRepo, mockEnvMngr)

		jwt, err := userUsecase.Login(req.Username, req.Password)
		assert.Error(t, err)
		assert.Nil(t, jwt)
	})
}
