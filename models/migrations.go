package models

import (
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var migrations = []*gormigrate.Migration{
	{
		ID: "01-schema-0010-01-channel",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&Channel{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("channels")
		},
	},
	{
		ID: "01-schema-0020-01-recording",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&Recording{},
				&RecordingFormat{})
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.Migrator().DropTable("recording_formats")
			if err != nil {
				return err
			}
			return tx.Migrator().DropTable("recordings")
		},
	},
}
