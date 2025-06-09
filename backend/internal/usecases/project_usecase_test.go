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
	"github.com/stretchr/testify/mock"
)

func TestCreateProject(t *testing.T) {
	t.Run("successfully create a project", func(t *testing.T) {
		// Arrange
		userId := uuid.NewString()
		createRequest := dtos.CreateProjectReq{
			Name:        "Test Project",
			Description: "A test project",
		}

		mockRepo := new(mocks.MockProjectRepository)

		mockRepo.On("Create", mock.MatchedBy(func(p *entities.Project) bool {
			return p.Name == createRequest.Name &&
				p.Description == createRequest.Description &&
				p.Id != "" &&
				p.UserId == userId
		})).Return(nil)

		service := usecases.NewProjectUsecase(mockRepo)

		// Act
		err := service.CreateProject(createRequest, userId)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("returns error when repo.Create fails", func(t *testing.T) {
		// Arrange
		userId := uuid.NewString()
		createRequest := dtos.CreateProjectReq{
			Name:        "Test Project",
			Description: "A test project",
		}

		mockRepo := new(mocks.MockProjectRepository)

		mockRepo.On("Create", mock.Anything).Return(errors.New("db error"))

		service := usecases.NewProjectUsecase(mockRepo)

		// Act
		err := service.CreateProject(createRequest, userId)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
	})
}
