package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Grand Piano: Kid be Kid
	testdataID := uuid.MustParse("84bc85bf-b304-4518-af66-afb17d8cbb54")
	testdataIDLang1 := uuid.MustParse("091a6639-a331-48ee-9b97-a8bf873c6da1")
	testdataIDSpeaker1 := uuid.MustParse("c32b348c-17da-48d2-af55-ab4f6f7a8036")
	testdataIDFormat1 := uuid.MustParse("a95f7f59-ad70-4902-bbb5-766a9f38d81d")

	// see stream 18

	/* WARNING unreleased:
	- Public
	- Private
	- Duration
	- Bytes
	*/

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-recording-23",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					CreatedAt: time.Date(2021, 7, 25, 20, 30, 0, 0, loc),
					Duration:  time.Hour,
					Public:    false,
					Listed:    false,
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker1,
							Name:    "Kid be Kid",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
					Lang:        "de",
					Title:       "Kid be Kid",
					Subtitle:    "Sonntag ab 20:30",
					Short:       `Gleich mit ihrem Debüt »Sold Out« hat sich **KID BE KID** schnurstracks auf die großen Festivals gebeamt und dort dem Publikum berauschende Konzerte geliefert.`,
					Long: `Den Abschluss unseres Konzertflügelfestivals legt **Kid be Kid** hin.

Gleich mit ihrem Debüt »Sold Out« hat sich **KID BE KID** schnurstracks auf die großen Festivals gebeamt und dort dem Publikum berauschende Konzerte geliefert. Ob beim Elbjazz in Hamburg, dem legendären Fusion Festival oder im bedeutenden Monument National im Rahmen des Festival International de Jazz de Montréal, eine der prominentesten Bühnen der Welt – **KID BE KID** fasziniert das Publikum mit bedingungsloser Hingabe und unglaublichem Können.
Mit ihrer neuen EP »Lovely Genders« etabliert **KID BE KID** ihren Platz zwischen den großen Neo Soul KünstlerInnen unserer Zeit. Klangliche Tiefe und inhaltliche Reife zeichnen ihre Songs aus. Ihr einzigartiges Skillset aus Beatboxing, Gesang, Klavier und Synthesizer trifft auf poetische Lyrics und rhythmische Virtuosität.
An **KID BE KID** ist einfach alles besonders. Vor ihr war eine Künstlerin kaum vorstellbar, die gleichzeitig vier Instrumente ganz ohne Loop Station spielt, dabei groovt und berührt bis zum Abwinken und oben drauf ihre Stimme in mehrere Töne spaltet. Sie verkörpert Coolness und Verletzlichkeit, Struktur und Freiheit, Hip Hop und Jazz gleichermaßen und lässt alles im Raum lebendig werden.`,
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
						Bytes:       0,
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
