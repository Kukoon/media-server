package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Grand Piano: Federico Albanese
	testdataStream := uuid.MustParse("bd2cb25e-6f17-453d-b947-370cd51beb18")
	testdataStreamLang1 := uuid.MustParse("6f188246-7b39-42e3-811c-a5828d01263b")
	testdataStreamSpeaker1 := uuid.MustParse("72496cb5-05cf-4982-8ec4-45666f3704e6")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-stream-14",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 7, 24, 18, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 7, 24, 17, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/bd2cb25e-6f17-453d-b947-370cd51beb18/poster.png",
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataStreamSpeaker1,
							Name:    "Federico Albanese",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStreamLang1,
					StreamID: testdataStream,
					Lang:     "de",
					Title:    "Federico Albanese",
					Subtitle: "Samstag ab 18:00",
					Short:    `**Federico Albanese** kombiniert in seinen Kompositionen minimalistische Pianomelodien mit Streicherarrangements und Elektronika.`,
					Long: `**Federico Albanese** kombiniert in seinen Kompositionen minimalistische Pianomelodien mit Streicherarrangements und Elektronika. Mit seiner Musik entführt der 1982 in Mailand geborene Wahlberliner in geradezu filmische Klangwelten, die sich aus Elementen der Klassik, Pop- und Ambientmusik speisen. Die vielfältigen musikalischen Einflüsse in seinen Kompositionen spiegeln zugleich seine große Liebe zur Musik.

Bereits im Kindesalter lernt **Federico Albanese** Klavier und Klarinette spielen, wendet sich als Teenager dem Gitarrenspiel zu und etabliert sich, nach einem Studium des Kontrabasses, als einer der führenden Protagonisten der Mailänder Untergrundszene.

Zusammen mit der Singer- und Songwriterin Jessica Einaudi, gründet er 2007 das Avantgarde-Duo »La Blanche Alchimie«, mit dem er drei Alben veröffentlicht und international Erfolge feiert. Auch als Komponist für Werbung, Film und Fernsehen ist Federico erfolgreich. Seine Scores sind z.B. in den internationalen Filmproduktionen »Shadows In The Distance« (2012), »Alles im grünen Bereich« (2014) oder dem ARTE-Dokumentarfilm »Cinema Perverso« (2015) zu hören.
**Federicos** Debüt als Solo-Pianist gibt er 2014 mit dem Album »The Houseboat and the Moon«. Mit »The Blue Hour« schließt er 2016 an und präsentiert klare, sehr dynamische und komplex arrangierte Kompositionen.`,
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
