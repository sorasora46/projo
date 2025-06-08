package usecases

import "github.com/sorasora46/projo/backend/internal/entities"

type ProjectUsecase interface {
	CreateProject(req any) error
	GetByProjectId(projectId string) (*entities.Project, error)
}

type ProjectService struct {
}

func NewProjectUsecase() ProjectUsecase {
	return &ProjectService{}
}

func (p *ProjectService) CreateProject(req any) error {
	return nil
}

func (p *ProjectService) GetByProjectId(projectId string) (*entities.Project, error) {
	return nil, nil
}
