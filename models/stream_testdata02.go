package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {

	testdataID := uuid.MustParse("0801a547-59f1-4a63-946f-2ab03f62e6ee")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0020-01-stream-02",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataID,
					ChannelID: TestChannelID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 3, 4, 18, 30, 0, 0, loc),
					ListenAt:  time.Date(2021, 3, 4, 18, 30, 0, 0, loc),
					Poster:    "https://c3woc.de/images/banner.jpg",
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Delete(&Stream{
					ID: testdataID,
				}).Error
			},
		},
	}...)
}
