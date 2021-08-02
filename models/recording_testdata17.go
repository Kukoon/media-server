package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Grand Piano: Niklas Paschburg
	testdataID := uuid.MustParse("54ff055b-5e46-4344-a43f-deb41c693045")
	testdataIDLang1 := uuid.MustParse("7b4275e0-2a11-4588-8264-51ff699d5868")
	testdataIDSpeaker1 := uuid.MustParse("6bd38420-f647-42ed-ba7c-99bd17b3cfe7")
	testdataIDFormat1 := uuid.MustParse("3509436c-284e-4ced-8959-fe27653fdc1b")

	// see stream 12

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-recording-17",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					CreatedAt: time.Date(2021, 7, 23, 20, 30, 0, 0, loc),
					Duration:  time.Hour + 18*time.Minute + 42*time.Second,
					Public:    true,
					Listed:    true,
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker1,
							Name:    "Niklas Paschburg",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
					Lang:        "de",
					Title:       "Niklas Paschburg",
					Subtitle:    "23. Juli 2021, 20:30",
					Short:       `2018 legte **Niklas Paschburg** sein Debütalbum Oceanic vor. Niklas‘ musikalische Ideen setzen sich aus einer Kombination umfangreicher Live-Performance sowie Experimenten zusammen. Im Februar 2020 erschien sein zweites Album Svalbard auf dem Label !7K.`,
					Long: `2018 legte **Niklas Paschburg** sein Debütalbum Oceanic vor. Niklas‘ musikalische Ideen setzen sich aus einer Kombination umfangreicher Live-Performance sowie Experimenten zusammen. Im Februar 2020 erschien sein zweites Album Svalbard auf dem Label !7K.

Niklas’ Gebrauch des Klaviers, elektronischer Elemente (Synth und Computer) sowie das Klavierakkordeon seines Großvaters ermöglichen es ihm, mit einer großen Bandbreite musikalischer Traditionen und Sprachen zu kommunizieren. Paschburgs Musik ist einzigartig in ihrer Fähigkeit, sowohl melancholisch als auch positiv zu sein — eine Umarmung, die Ängste abbaut und zur Meditation anregt, während sie die Hörer*in gleichzeitig dazu bringt, sich zur Musik zu bewegen, zu tanzen und zu rennen.
`,
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
						Bytes:       1612645697,
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
