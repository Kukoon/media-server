package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Tony Mahoni
	testdataID := uuid.MustParse("fb84591b-1ce5-4baf-8518-ccb38ed29b35")
	testdataIDLang1 := uuid.MustParse("efdfd37d-7f2b-4c90-ba5a-bec0a81920d6")
	testdataIDSpeaker1 := uuid.MustParse("13e32e59-43b5-4023-ba5f-0c1bdbc36a0f")
	testdataIDFormat1 := uuid.MustParse("9b271f38-556e-4f04-96b2-00f6647744f6")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-recording-22",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					CreatedAt: time.Date(2021, 8, 5, 17, 0, 0, 0, loc),
					Duration:  18*time.Minute + time.Second,
					Public:    true,
					Listed:    true,
					Tags: []*Tag{
						{ID: TestTagAusstellungID},
						{ID: TestTagInterviewID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker1,
							Name:    "Tony Mahoni",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
					Lang:        "de",
					Title:       "Tony Mahoni – »Vieja muito, vieja pouco«",
					Subtitle:    "Ein Gespräch mit zeit][raum zur Onlineausstellung",
					Short:       `**Tony Mahoni** malt seit seiner Kindheit gern Tiere und Pflanzen. Später hat er es einige Jahre fast ausschließlich auf Steine und Felsen abgesehen, bis dieses Motiv wiederum von dem des Wassers in Form von Wellen abgelöst wurde.`,
					Long:        `**Tony Mahoni** malt seit seiner Kindheit gern Tiere und Pflanzen. Später hat er es einige Jahre fast ausschließlich auf Steine und Felsen abgesehen, bis dieses Motiv wiederum von dem des Wassers in Form von Wellen abgelöst wurde. Nach einigen Jahren der künstlerischen Ungewissheit und der damit verbundenen Farblosigkeit hat er seinen Lebenslauf gefälscht, um nun, mit Beginn der gesellschaftlichen Krise wieder in den Farbbeutel zu greifen. Tatsächlich wurde der Zeitpunkt des beginnenden Lockdowns Anfang des Jahres 2020 zur Geburt seiner aktuellen Formen- und Farbsprache. Nach einem vierwöchigen künstlerischen Aufenthalt in Porto, sowie einer Kunstresidenz in Antwerpen, ist er nun erstmal wieder in Bremen zu Hause und nutzt die Chance seine Ölkreiden im Kukoon auszustellen.`,
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&[]*RecordingFormat{
					{
						ID:          testdataIDFormat1,
						RecordingID: testdataID,
						Lang:        "de",
						Quality:     0,
						IsVideo:     true,
						URL:         "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/video_best.mp4",
						Bytes:       424034431,
						Resolution:  "1920x1080",
					},
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Delete(&Recording{
					ID: testdataID,
				}).Error
			},
		},
	}...)
}
