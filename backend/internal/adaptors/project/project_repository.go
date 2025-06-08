package project

import (
	"github.com/sorasora46/projo/backend/internal/adaptors/interfaces"
	"github.com/sorasora46/projo/backend/internal/entities"
	"gorm.io/gorm"
)

type ProjectRepositoryImpl struct {
	db *gorm.DB
}

func NewProjectRepository() interfaces.ProjectRepository {
	return &ProjectRepositoryImpl{}
}

func (p *ProjectRepositoryImpl) Create(newProject entities.Project) error {
	return nil
}

func (p *ProjectRepositoryImpl) GetByProjectId(projectId string) (*entities.Project, error) {
	return nil, nil
}
