package interfaces

import "github.com/sorasora46/projo/backend/internal/entities"

type ProjectRepository interface {
	Create(newProject *entities.Project) error
	GetByProjectId(projectId string) (*entities.Project, error)
	GetAllProjects(userId string) ([]entities.Project, error)
}
