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
		ID: "01-schema-0017-01-event",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&Event{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("events")
		},
	},
	{
		ID: "01-schema-0018-01-tag",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&Tag{},
				&TagLang{})
		},
		Rollback: func(tx *gorm.DB) error {
			if err := tx.Migrator().DropTable("tag_langs"); err != nil {
				return err
			}
			return tx.Migrator().DropTable("tags")
		},
	},
	{
		ID: "01-schema-0010-01-speaker",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&Speaker{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("speakers")
		},
	},
	{
		ID: "01-schema-0020-01-recording",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&Recording{},
				&RecordingLang{},
				&RecordingFormat{})
		},
		Rollback: func(tx *gorm.DB) error {
			if err := tx.Migrator().DropTable("recording_formats"); err != nil {
				return err
			}
			if err := tx.Migrator().DropTable("recording_langs"); err != nil {
				return err
			}
			return tx.Migrator().DropTable("recordings")
		},
	},
	{
		ID: "01-schema-0030-01-stream",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&Stream{}, &StreamLang{})
		},
		Rollback: func(tx *gorm.DB) error {
			if err := tx.Migrator().DropTable(Stream{}.TableName()); err != nil {
				return err
			}
			return tx.Migrator().DropTable(Stream{}.TableName())
		},
	},
}
