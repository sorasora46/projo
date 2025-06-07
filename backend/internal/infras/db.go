package infras

import (
	"fmt"

	"github.com/sorasora46/projo/backend/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	InitDB(dsn string) error
	GetDBInstance() (*gorm.DB, error)
	migrateSchemas() error
}

type GormDatabase struct {
	db *gorm.DB
}

func NewDatabase() Database {
	return &GormDatabase{}
}

func (d *GormDatabase) InitDB(dsn string) error {
	var err error
	d.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := d.migrateSchemas(); err != nil {
		return nil
	}
	return nil
}

func (d *GormDatabase) GetDBInstance() (*gorm.DB, error) {
	if d.db == nil {
		return nil, fmt.Errorf("database instance is nil")
	}
	return d.db, nil
}

func (d *GormDatabase) migrateSchemas() error {
	if err := d.db.AutoMigrate(&entities.User{}); err != nil {
		return err
	}
	return nil
}
