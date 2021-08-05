package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Grand Piano: KONDRASCHEWA/CHICA
	testdataID := uuid.MustParse("9f8b59db-34e1-47c1-93ea-e173f5af3b96")
	testdataIDLang1 := uuid.MustParse("ebcaee38-56b9-4c70-9616-f88887bcfee4")
	testdataIDSpeaker1 := uuid.MustParse("3a875769-4b33-4fe7-af36-167c51fa510f")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0020-01-stream-13",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 7, 24, 16, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 7, 24, 12, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker1,
							Name:    "KONDRASCHEWA/CHICA",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataIDLang1,
					StreamID: testdataID,
					Lang:     "de",
					Title:    "KONDRASCHEWA/CHICA",
					Subtitle: "Samstag ab 16:00",
					Short:    `**KONDRASCHEWA/CHICA** bringen mit risikoreichem Verve und fast popmusikalisch anmutendem völlig unakademischen Appeal feinst gewobene Interpretationen von anspruchsvoller zeitgenössischer Klaviermusik für vier Hände und für zwei Klaviere zu Gehör.`,
					Long:     `**KONDRASCHEWA/CHICA** bringen mit risikoreichem Verve und fast popmusikalisch anmutendem völlig unakademischen Appeal feinst gewobene Interpretationen von anspruchsvoller zeitgenössischer Klaviermusik für vier Hände und für zwei Klaviere zu Gehör. Marina und Sebastian, aus der Ukraine/Deutschland und Kolumbien stammend, vermitteln in intimer Konzertatmosphäre die besondere Nähe dieser speziellen Art der Kammermusik. Das Programm sowie die Frische und Intensität der Interpretation sind absolut hörenswert, auch für Hörer, die bisher mit sogenannter “klassischer” Klaviermusik nicht soviel anzufangen wissen.`,
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
