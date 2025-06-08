package entities

type Project struct {
	Id     string `gorm:"primaryKey"`
	Name   string `gorm:"not null"`
	UserId string
	Model
}
