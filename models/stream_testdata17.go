package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Prozess gegen den Halle-Attentäter
	testdataID := uuid.MustParse("54c54061-d0f9-4852-ad98-fdea350ef174")
	testdataIDLang1 := uuid.MustParse("68a59a91-3df7-4a35-828b-e2f736e6cd7a")
	testdataIDSpeaker1 := uuid.MustParse("eac66a41-12ef-40d6-9eca-5d7063de1d8a")
	testdataIDSpeaker2 := uuid.MustParse("ea1a81fa-dce0-4026-8d83-1f069c0eb8ad")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0020-01-stream-17",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 7, 29, 19, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 7, 27, 0, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker1,
							Name:    "Naomi Henkel-Gümbel",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker2,
							Name:    "Dr. Kati Lang",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataIDLang1,
					StreamID: testdataID,
					Lang:     "de",
					Title:    "Prozess gegen den Halle-Attentäter",
					Subtitle: "Gespräch mit Nebenklägerin Naomi Henkel-Gümbel und Nebenklageanwältin Dr. Kati Lang",
					Short:    `Am 9. Oktober 2019 verübte ein Neonazi einen rechtsterroristischen Anschlag in Halle(Saale).`,
					Long: `Am 9. Oktober 2019 verübte ein Neonazi einen rechtsterroristischen Anschlag in Halle(Saale). Am höchsten jüdischen Feiertag, Yom Kippur, versuchte er in die Synagoge einzudringen, um die dort betenden Jüdinnen und Juden zu töten. Nachdem dieser Versuch gescheitert war, wich er auf einen nahegelegenen Imbiss, den Kiezdöner aus, um dort vermeintliche Muslime zu töten. Während des Anschlags tötete der Neonazi zwei Menschen und fuhr auf seiner Flucht eine weitere Person aus rassistischem Motiv mit Tötungsabsicht an.
Vom Juli bis Dezember 2020 verlief gegen den Attentäter der Gerichtsprozess, der für viele Facetten der deutschen Gesellschaft als beispielhaft gesehen werden kann. Dazu gehören Aspekte wie die Stilisierung des Attentäters als Einzeltäter, das hohe Interesse der Medien an der privaten Geschichte des Attentäters und die damit gebotene Plattform, weiterhin die zunächst ausbleibende Zulassung der Synagogenbesucher*innen in die Nebenklage, und schließlich die fehlende Anerkennung von versuchtem Mord in zwei eindeutigen Fällen.
Was es mit diesen und weiteren Aspekten auf sich hat, wollen wir im Gespräch mit Naomi Henkel-Gümbel und Dr. Kati Lang auf den Grund gehen. Aus persönlicher Perspektive als Nebenklägerin und Anwältin werden die beiden Gäste über ihre Erfahrungen und Beobachtungen berichten und einen Einblick geben, was der Prozess für den deutschen Staat und die deutsche Gesellschaft bedeutet hat.

Eine Veranstaltung des Jungen Forum der DIG Bremen, der Falken Bremen, der DIG Bremen/Unterweser, mbt gegen Rechtsextremismus und der Naturfreundejugend Bremen.`,
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
