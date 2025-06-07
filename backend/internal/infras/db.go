package infras

import (
	"fmt"
	"log"

	"github.com/sorasora46/projo/backend/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	InitDB(dsn string)
	GetDBInstance() (*gorm.DB, error)
	migrateSchemas() error
}

type GormDatabase struct {
	db *gorm.DB
}

func NewDatabase() Database {
	return &GormDatabase{}
}

func (d *GormDatabase) InitDB(dsn string) {
	var err error
	d.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := d.migrateSchemas(); err != nil {
		log.Fatal(err)
	}
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
