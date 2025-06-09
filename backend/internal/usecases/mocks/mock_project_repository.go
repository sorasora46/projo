package mocks

import (
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
