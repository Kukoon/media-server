package models

import (
	"time"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var testdata = []*gormigrate.Migration{
	{
		ID: "10-data-0020-01-recording",
		Migrate: func(tx *gorm.DB) error {
			err := tx.Create(&Recording{
				ID:         uuid.MustParse("dffe2c0e-3713-4399-8ee2-279becbbb06e"),
				CommonName: "2020-12-polizeigewalt",
				Duration:   time.Hour,
			}).Error
			if err != nil {
				return err
			}
			return tx.Create(&RecordingFormat{
				ID:          uuid.MustParse("3a4f9157-65bf-4d15-a82b-1cd9295d07e0"),
				RecordingID: uuid.MustParse("dffe2c0e-3713-4399-8ee2-279becbbb06e"),
				IsVideo:     true,
				URL:         "https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee_20201211_165251.mp4",
				Resolution:  "1920x1080",
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.Delete(&RecordingFormat{
				ID: uuid.MustParse("3a4f9157-65bf-4d15-a82b-1cd9295d07e0"),
			}).Error
			if err != nil {
				return err
			}
			return tx.Delete(&Recording{
				ID: uuid.MustParse("dffe2c0e-3713-4399-8ee2-279becbbb06e"),
			}).Error
		},
	},
}
