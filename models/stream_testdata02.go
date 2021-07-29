package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {

	testdataID := uuid.MustParse("0801a547-59f1-4a63-946f-2ab03f62e6ee")
	testdataIDLang1 := uuid.MustParse("50388781-0c42-434e-ae99-170e4c1c0bcf")

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
				if err := tx.Create(&StreamLang{
					ID:       testdataIDLang1,
					StreamID: testdataID,
					Lang:     "de",
					Title:    "C3 Waffel Operation Center (C3WOC)",
					Subtitle: "Seid dort nett zu einander und freut euch auf viele Videos rund um die Waffel!",
					Short:    `Für den kleinen Hunger empfehlen wir:`,
					Long: `Für den kleinen Hunger empfehlen wir:

4 Eier
200g Zucker
250g Butter
15g Vanillezucker
10g Backpulver
200ml Vollmilch
400g Mehl
120ml Mate
etwas Rum
eine prise Salz`,
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
