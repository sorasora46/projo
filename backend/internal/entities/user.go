package entities

type User struct {
	Id             string `gorm:"primaryKey"`
	FirstName      string `gorm:"not null"`
	LastName       string `gorm:"not null"`
	Username       string `gorm:"unique;not null"`
	Email          string `gorm:"unique;not null"`
	HashedPassword []byte `gorm:"not null"`
}
