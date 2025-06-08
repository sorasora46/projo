package entities

type ProjectTask struct {
	Id          string `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	status      string `gorm:"not null"` // TODO, DOING, DONE
	ProjectId   string
	Model
}
