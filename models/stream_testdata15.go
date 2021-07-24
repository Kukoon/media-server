package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Grand Piano: Clara Vetter
	testdataStream := uuid.MustParse("28cf0c84-07fe-4421-b2dc-7ad6c2551689")
	testdataStreamLang1 := uuid.MustParse("5961b97f-c19d-40ec-8a6a-562be6b9b083")
	testdataStreamSpeaker1 := testdataStream8Speaker3

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-stream-15",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 7, 24, 20, 30, 0, 0, loc),
					ListenAt:  time.Date(2021, 7, 24, 20, 00, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/28cf0c84-07fe-4421-b2dc-7ad6c2551689/poster.png",
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataStreamSpeaker1,
							Name:    "Clara Vetter",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStreamLang1,
					StreamID: testdataStream,
					Lang:     "de",
					Title:    "Clara Vetter",
					Subtitle: "Samstag ab 20:30",
					Short:    `Pianistin und Komponistin **Clara Vetter** (*1996) widmet sich beinahe lebenslang der Musik und anderen kreativen Schaffensprozessen.`,
					Long:     `Pianistin und Komponistin **Clara Vetter** (*1996) widmet sich beinahe lebenslang der Musik und anderen kreativen Schaffensprozessen. An der Hochschule für Musik und Darstellende Kunst Stuttgart absolvierte sie 2018 ihr Bachelorstudium in Jazz Klavier und wurde im selben Jahr mit dem Stuttgarter Steinway&Sons Förderpreis ausgezeichnet. Im Juni 2021 schloss sie ihr Masterstudium in Music Performance am Kopenhagener Rytmisk Musikkonservatorium ab. Natürlichkeit - "organisch sein" - ist, was sie an Kunst am meisten fasziniert und somit geht ihre Musik stets einer komplexen, dennoch fließenden Lebendigkeit entgegen.`,
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
