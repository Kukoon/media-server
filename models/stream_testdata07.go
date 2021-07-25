package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Grand Piano: Martin Kohlstedt
	testdataStream := uuid.MustParse("7fc21416-5d68-4ecf-bd4b-e8a89f7c31f2")
	testdataStreamLang1 := uuid.MustParse("45bb0727-4812-40a2-8ac2-12422dfb42f1")
	testdataStreamSpeaker1 := uuid.MustParse("976010a0-c19f-4d22-a4d6-9553b460adfe")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-stream-07",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 7, 21, 0, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 7, 21, 0, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataStream.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataStream.String() + "/preview.webp",
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataStreamSpeaker1,
							Name:    "Martin Kohlstedt",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStreamLang1,
					StreamID: testdataStream,
					Lang:     "de",
					Title:    "Martin Kohlstedt",
					Subtitle: "Mittwoch ab 18:00",
					Short:    `Kein geringerer als **Martin Kohlstedt** wird den wunderbaren Auftakt machen und den Park auf die kommenden Konzerte perfekt einstimmen.`,
					Long: `Kein geringerer als **Martin Kohlstedt** wird den wunderbaren Auftakt machen und den Park auf die kommenden Konzerte perfekt einstimmen.
Das neue Martin Kohlstedt Album »FLUR« erschien im November 2020 auf Warner Classics. Das besondere Setup aus Flügel, Synthesizern und Electronika, kombiniert mit Kohlstedts Ansatz jedes Konzert von Grund auf neu zu verhandeln macht seine Konzerte zu einem Erlebnis.

**Martin Kohlstedt** lebt und arbeitet in Weimar. Seine bisherigen Alben TAG, NACHT, STROM, STRÖME und deren Begleiter in Form von Reworks erhielten internationale Anerkennung und führten den Komponisten und Pianisten auf Konzertreisen in der ganzen Welt.

**Kohlstedt** bezeichnet seine Art des Arbeitens als modulares Komponieren, die Stücke sind ständig in Bewegung und folgen auch im Konzert keiner festgelegten Form. Improvisation ist zwingend Teil des Schaffens des 1988 geborenen Musikers, ebenso wie Augenhöhe mit dem Publikum, der Mut zu Scheitern und die Interaktion mit Raum, Menschen und Kontext.`,
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