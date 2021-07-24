package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Grand Piano: KONDRASCHEWA/CHICA
	testdataStream := uuid.MustParse("9f8b59db-34e1-47c1-93ea-e173f5af3b96")
	testdataStreamLang1 := uuid.MustParse("ebcaee38-56b9-4c70-9616-f88887bcfee4")
	testdataStreamSpeaker1 := uuid.MustParse("3a875769-4b33-4fe7-af36-167c51fa510f")
	testdataStreamSpeaker2 := uuid.MustParse("4fef5401-42c3-45a7-9403-1a40a3a0d946")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-stream-13",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 7, 24, 16, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 7, 24, 12, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/9f8b59db-34e1-47c1-93ea-e173f5af3b96/poster.png",
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataStreamSpeaker1,
							Name:    "MARINA KONDRASCHEWA",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataStreamSpeaker2,
							Name:    "SEBASTIAN CHICA VILLA",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStreamLang1,
					StreamID: testdataStream,
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
					ID: testdataStream,
				}).Error
			},
		},
	}...)
}
