package database

import (
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB                *gorm.DB
	Connection        string          `toml:"connection"`
	Debug             bool            `toml:"debug"`
	Testdata          bool            `toml:"testdata"`
	LogLevel          logger.LogLevel `toml:"log_level"`
	migrations        map[string]*gormigrate.Migration
	migrationTestdata map[string]*gormigrate.Migration
}

func (config *Database) Run() error {
	db, err := gorm.Open(postgres.Open(config.Connection), &gorm.Config{
		Logger: logger.Default.LogMode(config.LogLevel),
	})
	if err != nil {
		return err
	}
	db.Debug().Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if config.Debug {
		db = db.Debug()
	}

	config.DB = db
	if err = config.migrate(config.Testdata); err != nil {
		return err
	}
	return nil
}

func (config *Database) Status() error {
	sqlDB, err := config.DB.DB()
	if err != nil {
		return err

	}
	if err = sqlDB.Ping(); err != nil {
		return err
	}
	return nil
}
