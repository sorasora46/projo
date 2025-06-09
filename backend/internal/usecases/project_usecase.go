package usecases

import (
	"github.com/google/uuid"
	"github.com/sorasora46/projo/backend/internal/adaptors/interfaces"
	"github.com/sorasora46/projo/backend/internal/dtos"
	"github.com/sorasora46/projo/backend/internal/entities"
)

type ProjectUsecase interface {
	CreateProject(req dtos.CreateProjectReq, userId string) error
	GetByProjectId(projectId string) (*entities.Project, error)
}

type ProjectService struct {
	repo interfaces.ProjectRepository
}

func NewProjectUsecase(repo interfaces.ProjectRepository) ProjectUsecase {
	return &ProjectService{repo: repo}
}

func (p *ProjectService) CreateProject(req dtos.CreateProjectReq, userId string) error {
	newProject := &entities.Project{
		Id:          uuid.NewString(),
		Name:        req.Name,
		Description: req.Description,
		UserId:      userId,
	}
	if err := p.repo.Create(newProject); err != nil {
		return err
	}
	return nil
}

func (p *ProjectService) GetByProjectId(projectId string) (*entities.Project, error) {
	return nil, nil
}
