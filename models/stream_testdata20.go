package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Seyda Kurt - Radikale Zärtlichkeit
	testdataID := uuid.MustParse("a7816cbb-3fab-40e4-a753-a87d2439df7f")
	testdataIDLang1 := uuid.MustParse("a87b4871-36f4-47b5-b181-8e3d2525c8c0")
	testdataIDSpeaker1 := uuid.MustParse("7885aa1a-a564-4547-b3ba-ca0ec0cd7fa3")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0020-01-stream-20",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 9, 2, 18, 30, 0, 0, loc),
					ListenAt:  time.Date(2021, 0, 0, 0, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.jpg",
					Tags: []*Tag{
						{ID: TestTagBuchvorstellungID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker1,
							Name:    "Şeyda Kurt",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataIDLang1,
					StreamID: testdataID,
					Lang:     "de",
					Title:    "Radikale Zärtlichkeit",
					Subtitle: "Warum Liebe politisch ist",
					Short:    `Für Şeyda Kurt ist eine Ethik der radikalen Zärtlichkeit eine Ethik der Gerechtigkeit.`,
					Long: `Für Şeyda Kurt ist eine Ethik der radikalen Zärtlichkeit eine Ethik der Gerechtigkeit. Eine Ethik eines neuen Miteinanders, das sich gegen die gewaltvollen Logiken und ihren Anspruch auf Unumstößlichkeit stellt. Kurts Analyse und Kritik gründen auf einem Unbehagen mit und in den gegenwärtigen Verhältnissen. Zärtlichkeit ist dabei Ausgangspunkt und Ziel ihrer Arbeit. In ihrem Buch »Radikale Zärtlichkeit. Warum Liebe politisch ist« schaut Kurt aus postkolonialer, antikapitalistischer, feministischer und klassismuskritischer Perspektive (nicht nur) auf emotionale und romantische Selbstverständnisse. Diese sollten uns Anlass geben, in eine Krise der Wahrheiten zu geraten. Und Raum für Utopien entstehen zu lassen. Für beides nehmen wir uns in dieser Veranstaltung Zeit!
 
Eine Veranstaltung des __Queer Power Month Bremen__ in Kooperation mit dem __Kukoon__`,
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
