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

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-stream-12",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 7, 23, 20, 15, 0, 0, loc),
					ListenAt:  time.Date(2021, 7, 23, 20, 15, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
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
				if err := tx.Create(&StreamLang{
					ID:       testdataIDLang1,
					StreamID: testdataID,
					Lang:     "de",
					Title:    "Niklas Paschburg",
					Subtitle: "Freitag ab 20:30",
					Short:    `2018 legte **Niklas Paschburg** sein Debütalbum Oceanic vor. Niklas‘ musikalische Ideen setzen sich aus einer Kombination umfangreicher Live-Performance sowie Experimenten zusammen. Im Februar 2020 erschien sein zweites Album Svalbard auf dem Label !7K.`,
					Long: `2018 legte **Niklas Paschburg** sein Debütalbum Oceanic vor. Niklas‘ musikalische Ideen setzen sich aus einer Kombination umfangreicher Live-Performance sowie Experimenten zusammen. Im Februar 2020 erschien sein zweites Album Svalbard auf dem Label !7K.

Niklas’ Gebrauch des Klaviers, elektronischer Elemente (Synth und Computer) sowie das Klavierakkordeon seines Großvaters ermöglichen es ihm, mit einer großen Bandbreite musikalischer Traditionen und Sprachen zu kommunizieren. Paschburgs Musik ist einzigartig in ihrer Fähigkeit, sowohl melancholisch als auch positiv zu sein — eine Umarmung, die Ängste abbaut und zur Meditation anregt, während sie die Hörer*in gleichzeitig dazu bringt, sich zur Musik zu bewegen, zu tanzen und zu rennen.
`,
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
