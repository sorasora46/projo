package project

import (
	"github.com/sorasora46/projo/backend/internal/adaptors/interfaces"
	"github.com/sorasora46/projo/backend/internal/entities"
	"gorm.io/gorm"
)

type ProjectRepositoryImpl struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) interfaces.ProjectRepository {
	return &ProjectRepositoryImpl{db: db}
}

func (p *ProjectRepositoryImpl) Create(newProject *entities.Project) error {
	transaction := p.db.Create(newProject)
	if transaction.Error != nil {
		return transaction.Error
	}
	return nil
}

func (p *ProjectRepositoryImpl) GetByProjectId(projectId string) (*entities.Project, error) {
	var project entities.Project
	transaction := p.db.First(&project)
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	return &project, nil
}

func (p *ProjectRepositoryImpl) GetAllProjects(userId string) ([]entities.Project, error) {
	var projects []entities.Project
	transaction := p.db.Where("user_id = ?", userId).Find(&projects)
	if transaction.Error != nil {
		return nil, transaction.Error
	}

	return projects, nil
}

func (p *ProjectRepositoryImpl) DeleteByProjectId(projectId string) error {
	project := &entities.Project{Id: projectId}
	transaction := p.db.Delete(project)
	if transaction.Error != nil {
		return transaction.Error
	}
	return nil
}

func (p *ProjectRepositoryImpl) CheckIfProjectExistById(projectId string) (bool, error) {
	var count int64
	transaction := p.db.Where("id = ?", projectId).Find(&entities.Project{}).Count(&count)
	if transaction.Error != nil {
		return false, transaction.Error
	}
	return count == 1, nil
}
