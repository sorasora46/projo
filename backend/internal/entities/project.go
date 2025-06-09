package entities

type Project struct {
	Id           string `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Description  string
	UserId       string        `gorm:"not null"`
	ProjectTasks []ProjectTask `gorm:"constraint:OnDelete:CASCADE"`
	Model
}
