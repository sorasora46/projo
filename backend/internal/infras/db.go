package infras

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sorasora46/projo/backend/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	var err error
	d.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: dbLogger})
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
		log.Fatal("error migrating user entity")
		return err
	}
	if err := d.db.AutoMigrate(&entities.Project{}); err != nil {
		log.Fatal("error migrating project entity")
		return err
	}
	if err := d.db.AutoMigrate(&entities.ProjectTask{}); err != nil {
		log.Fatal("error migrating project_task entity")
		return err
	}
	return nil
}
