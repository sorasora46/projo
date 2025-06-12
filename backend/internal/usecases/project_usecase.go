package usecases

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sorasora46/projo/backend/internal/adaptors/interfaces"
	"github.com/sorasora46/projo/backend/internal/dtos"
	"github.com/sorasora46/projo/backend/internal/entities"
)

type ProjectUsecase interface {
	CreateProject(req dtos.CreateProjectReq, userId string) error
	GetByProjectId(projectId string) (*entities.Project, error)
	GetAllProjects(userId string) ([]entities.Project, error)
	DeleteByProjectId(projectId string) error
	UpdateProject(req dtos.UpdateProjectReq, projectId string) error
	ProjectShouldExist(projectId string) error
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
	project, err := p.repo.GetByProjectId(projectId)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (p *ProjectService) GetAllProjects(userId string) ([]entities.Project, error) {
	projects, err := p.repo.GetAllProjects(userId)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (p *ProjectService) DeleteByProjectId(projectId string) error {
	if err := p.ProjectShouldExist(projectId); err != nil {
		return err
	}
	if err := p.repo.DeleteByProjectId(projectId); err != nil {
		return err
	}
	return nil
}

func (p *ProjectService) UpdateProject(req dtos.UpdateProjectReq, projectId string) error {
	if err := p.ProjectShouldExist(projectId); err != nil {
		return err
	}
	if err := p.repo.UpdateProject(req, projectId); err != nil {
		return err
	}

	return nil
}

func (p *ProjectService) ProjectShouldExist(projectId string) error {
	isExist, err := p.repo.CheckIfProjectExistById(projectId)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New("project not exist")
	}
	return nil
}
