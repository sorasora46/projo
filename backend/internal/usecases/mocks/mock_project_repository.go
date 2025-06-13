package mocks

import (
	"github.com/sorasora46/projo/backend/internal/dtos/req"
	"github.com/sorasora46/projo/backend/internal/entities"
	"github.com/stretchr/testify/mock"
)

type MockProjectRepository struct {
	mock.Mock
}

func (m *MockProjectRepository) Create(newProject *entities.Project) error {
	args := m.Called(newProject)
	return args.Error(0)
}

func (m *MockProjectRepository) GetByProjectId(projectId string) (*entities.Project, error) {
	args := m.Called(projectId)
	if proj := args.Get(0); proj != nil {
		return proj.(*entities.Project), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockProjectRepository) GetAllProjects(userId string) ([]entities.Project, error) {
	args := m.Called(userId)
	if projects := args.Get(0); projects != nil {
		return projects.([]entities.Project), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockProjectRepository) DeleteByProjectId(projectId string) error {
	args := m.Called(projectId)
	return args.Error(0)
}

func (m *MockProjectRepository) CheckIfProjectExistById(projectId string) (bool, error) {
	args := m.Called(projectId)
	return args.Bool(0), args.Error(1)
}

func (m *MockProjectRepository) UpdateProject(req req.UpdateProjectReq, projectId string) error {
	args := m.Called(req, projectId)
	return args.Error(0)
}
