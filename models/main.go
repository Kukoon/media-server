package models

import (
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB         *gorm.DB
	Connection string          `toml:"connection"`
	Debug      bool            `toml:"debug"`
	Testdata   bool            `toml:"testdata"`
	LogLevel   logger.LogLevel `toml:"log_level"`
}

func (config *Database) Run() error {
	db, err := gorm.Open(postgres.Open(config.Connection), &gorm.Config{
		Logger: logger.Default.LogMode(config.LogLevel),
	})
	if err != nil {
		return err
	}
	if config.Debug {
		db = db.Debug()
	}

	mig := migrations
	if config.Testdata {
		mig = append(mig, testdata...)
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, mig)

	if err = m.Migrate(); err != nil {
		return nil
	}
	config.DB = db
	return nil
}
