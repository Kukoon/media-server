package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Grand Piano: Federico Albanese
	testdataID := uuid.MustParse("bd2cb25e-6f17-453d-b947-370cd51beb18")
	testdataIDLang1 := uuid.MustParse("6f188246-7b39-42e3-811c-a5828d01263b")
	testdataIDSpeaker1 := uuid.MustParse("72496cb5-05cf-4982-8ec4-45666f3704e6")
	testdataIDFormat1 := uuid.MustParse("3aa420ba-7eee-4f86-9a69-ea6e363733f5")

	// see stream 14

	/* WARNING unreleased:
	- Public
	- Private
	- Duration
	- Bytes
	*/

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-recording-19",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					CreatedAt: time.Date(2021, 7, 24, 18, 0, 0, 0, loc),
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
							Name:    "Federico Albanese",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
					Lang:        "de",
					Title:       "Federico Albanese",
					Subtitle:    "Samstag ab 18:00",
					Short:       `**Federico Albanese** kombiniert in seinen Kompositionen minimalistische Pianomelodien mit Streicherarrangements und Elektronika.`,
					Long: `**Federico Albanese** kombiniert in seinen Kompositionen minimalistische Pianomelodien mit Streicherarrangements und Elektronika. Mit seiner Musik entführt der 1982 in Mailand geborene Wahlberliner in geradezu filmische Klangwelten, die sich aus Elementen der Klassik, Pop- und Ambientmusik speisen. Die vielfältigen musikalischen Einflüsse in seinen Kompositionen spiegeln zugleich seine große Liebe zur Musik.

Bereits im Kindesalter lernt **Federico Albanese** Klavier und Klarinette spielen, wendet sich als Teenager dem Gitarrenspiel zu und etabliert sich, nach einem Studium des Kontrabasses, als einer der führenden Protagonisten der Mailänder Untergrundszene.

Zusammen mit der Singer- und Songwriterin Jessica Einaudi, gründet er 2007 das Avantgarde-Duo »La Blanche Alchimie«, mit dem er drei Alben veröffentlicht und international Erfolge feiert. Auch als Komponist für Werbung, Film und Fernsehen ist Federico erfolgreich. Seine Scores sind z.B. in den internationalen Filmproduktionen »Shadows In The Distance« (2012), »Alles im grünen Bereich« (2014) oder dem ARTE-Dokumentarfilm »Cinema Perverso« (2015) zu hören.
**Federicos** Debüt als Solo-Pianist gibt er 2014 mit dem Album »The Houseboat and the Moon«. Mit »The Blue Hour« schließt er 2016 an und präsentiert klare, sehr dynamische und komplex arrangierte Kompositionen.`,
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
