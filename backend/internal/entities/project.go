package entities

type Project struct {
	Id           string        `gorm:"primaryKey" json:"id"`
	Name         string        `gorm:"not null" json:"name"`
	Description  string        `json:"description"`
	UserId       string        `gorm:"not null" json:"userId"`
	ProjectTasks []ProjectTask `gorm:"constraint:OnDelete:CASCADE" json:"projectTasks"`
	Model
}
