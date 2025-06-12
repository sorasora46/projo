package usecases_test

import (
	"errors"
	"testing"
	"time"

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

func TestGetByProjectId(t *testing.T) {
	t.Run("successfully get project by projectId", func(t *testing.T) {
		// Arrange
		projectId := uuid.NewString()
		userId := uuid.NewString()
		now := time.Now()
		expectedProject := &entities.Project{
			Id:          projectId,
			Name:        "project name",
			Description: "project description",
			UserId:      userId,
			Model: entities.Model{
				CreatedAt: now,
				UpdatedAt: now,
			},
		}

		mockRepo := new(mocks.MockProjectRepository)

		mockRepo.On("GetByProjectId", mock.AnythingOfType("string")).Return(expectedProject, nil)

		service := usecases.NewProjectUsecase(mockRepo)

		// Act
		project, err := service.GetByProjectId(projectId)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedProject, project)
		mockRepo.AssertExpectations(t)
	})

	t.Run("returns error when repo.GetByProjectId fails", func(t *testing.T) {
		// Arrange
		projectId := uuid.NewString()

		mockRepo := new(mocks.MockProjectRepository)

		mockRepo.On("GetByProjectId", mock.AnythingOfType("string")).Return(nil, errors.New("db error"))

		service := usecases.NewProjectUsecase(mockRepo)

		// Act
		project, err := service.GetByProjectId(projectId)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, project)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetAllProjects(t *testing.T) {
	t.Run("successfully get all projects", func(t *testing.T) {
		// Arrange
		userId := uuid.NewString()
		projectId := uuid.NewString()
		expectedRes := []entities.Project{
			{
				UserId:      userId,
				Id:          projectId,
				Name:        "project 1",
				Description: "description 1",
			},
		}
		mockRepo := new(mocks.MockProjectRepository)
		mockRepo.On("GetAllProjects", mock.AnythingOfType("string")).Return(expectedRes, nil)

		service := usecases.NewProjectUsecase(mockRepo)

		// Act
		projects, err := service.GetAllProjects(userId)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, len(expectedRes), len(projects))
		assert.Equal(t, expectedRes[0].Id, projects[0].Id)
		assert.Equal(t, expectedRes[0].Name, projects[0].Name)
		assert.Equal(t, expectedRes[0].Description, projects[0].Description)
		assert.Equal(t, expectedRes[0].UserId, projects[0].UserId)
		mockRepo.AssertExpectations(t)
	})

	t.Run("returns error when repo.GetAllProjects fails", func(t *testing.T) {
		// Arrange
		userId := uuid.NewString()
		mockRepo := new(mocks.MockProjectRepository)
		mockRepo.On("GetAllProjects", mock.AnythingOfType("string")).Return(nil, errors.New("db failure"))

		service := usecases.NewProjectUsecase(mockRepo)

		// Act
		projects, err := service.GetAllProjects(userId)
		// Assert
		assert.Error(t, err)
		assert.Nil(t, projects)
		assert.Equal(t, "db failure", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateProject(t *testing.T) {
	t.Run("successfully update project", func(t *testing.T) {
		// Arrange
		projectId := uuid.NewString()
		updateProj := dtos.UpdateProjectReq{
			Name:        "project edited",
			Description: "description edited",
		}
		mockRepo := new(mocks.MockProjectRepository)
		mockRepo.On("CheckIfProjectExistById", mock.AnythingOfType("string")).Return(true, nil)
		mockRepo.On("UpdateProject", mock.AnythingOfType("dtos.UpdateProjectReq"), mock.AnythingOfType("string")).Return(nil)

		service := usecases.NewProjectUsecase(mockRepo)

		// Act
		err := service.UpdateProject(updateProj, projectId)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("returns error when repo.CheckIfProjectExistById fails", func(t *testing.T) {
		// Arrange
		projectId := uuid.NewString()
		updateProj := dtos.UpdateProjectReq{
			Name:        "project edited",
			Description: "description edited",
		}
		mockRepo := new(mocks.MockProjectRepository)
		mockRepo.On("CheckIfProjectExistById", mock.AnythingOfType("string")).Return(false, errors.New("db failure"))

		service := usecases.NewProjectUsecase(mockRepo)

		// Act
		err := service.UpdateProject(updateProj, projectId)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "db failure", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("returns error when repo.CheckIfProjectExistById return false", func(t *testing.T) {
		// Arrange
		projectId := uuid.NewString()
		updateProj := dtos.UpdateProjectReq{
			Name:        "project edited",
			Description: "description edited",
		}
		mockRepo := new(mocks.MockProjectRepository)
		mockRepo.On("CheckIfProjectExistById", mock.AnythingOfType("string")).Return(false, nil)

		service := usecases.NewProjectUsecase(mockRepo)

		// Act
		err := service.UpdateProject(updateProj, projectId)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "project not exist", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("returns error when repo.UpdateProjet fails", func(t *testing.T) {
		// Arrange
		projectId := uuid.NewString()
		updateProj := dtos.UpdateProjectReq{
			Name:        "project edited",
			Description: "description edited",
		}
		mockRepo := new(mocks.MockProjectRepository)
		mockRepo.On("CheckIfProjectExistById", mock.AnythingOfType("string")).Return(true, nil)
		mockRepo.On("UpdateProject", mock.AnythingOfType("dtos.UpdateProjectReq"), mock.AnythingOfType("string")).Return(errors.New("db failure"))

		service := usecases.NewProjectUsecase(mockRepo)

		// Act
		err := service.UpdateProject(updateProj, projectId)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "db failure", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteByProjectId(t *testing.T) {
	t.Run("successfully delete project by Id", func(t *testing.T) {
		// Arrange
		projectId := uuid.NewString()
		mockRepo := new(mocks.MockProjectRepository)
		mockRepo.On("CheckIfProjectExistById", mock.AnythingOfType("string")).Return(true, nil)
		mockRepo.On("DeleteByProjectId", mock.AnythingOfType("string")).Return(nil)

		service := usecases.NewProjectUsecase(mockRepo)
		// Act
		err := service.DeleteByProjectId(projectId)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("returns error when repo.CheckIfProjectExist return false", func(t *testing.T) {
		// Arrange
		projectId := uuid.NewString()
		mockRepo := new(mocks.MockProjectRepository)
		mockRepo.On("CheckIfProjectExistById", mock.AnythingOfType("string")).Return(false, nil)

		service := usecases.NewProjectUsecase(mockRepo)
		// Act
		err := service.DeleteByProjectId(projectId)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "project not exist", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("returns error when repo.CheckIfProjectExist fails", func(t *testing.T) {
		// Arrange
		projectId := uuid.NewString()
		mockRepo := new(mocks.MockProjectRepository)
		mockRepo.On("CheckIfProjectExistById", mock.AnythingOfType("string")).Return(false, errors.New("db failure"))

		service := usecases.NewProjectUsecase(mockRepo)
		// Act
		err := service.DeleteByProjectId(projectId)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "db failure", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("returns error when repo.DeleteByProjectId fails", func(t *testing.T) {
		// Arrange
		projectId := uuid.NewString()
		mockRepo := new(mocks.MockProjectRepository)
		mockRepo.On("CheckIfProjectExistById", mock.AnythingOfType("string")).Return(true, nil)
		mockRepo.On("DeleteByProjectId", mock.AnythingOfType("string")).Return(errors.New("db failure"))

		service := usecases.NewProjectUsecase(mockRepo)
		// Act
		err := service.DeleteByProjectId(projectId)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "db failure", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
