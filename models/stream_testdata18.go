package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Grand Piano: Kid be Kid
	testdataStream := uuid.MustParse("84bc85bf-b304-4518-af66-afb17d8cbb54")
	testdataStreamLang1 := uuid.MustParse("091a6639-a331-48ee-9b97-a8bf873c6da1")
	testdataStreamSpeaker1 := uuid.MustParse("c32b348c-17da-48d2-af55-ab4f6f7a8036")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-stream-18",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 7, 25, 20, 30, 0, 0, loc),
					ListenAt:  time.Date(2021, 7, 25, 20, 00, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/84bc85bf-b304-4518-af66-afb17d8cbb54/poster.png",
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataStreamSpeaker1,
							Name:    "Kid be Kid",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStreamLang1,
					StreamID: testdataStream,
					Lang:     "de",
					Title:    "Kid be Kid",
					Subtitle: "Sonntag ab 20:30",
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
