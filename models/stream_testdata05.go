package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// see recording 10
	testdataID := uuid.MustParse("4fb029d6-063a-4302-9ae8-4c1c6a1542a5")
	testdataIDLang1 := uuid.MustParse("d5262bb7-378b-456f-9e91-34f63b174c48")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0020-01-stream-05",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID1,
					Chat:      true,
					Running:   true,
					StartAt:   time.Date(2021, 5, 5, 0, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 5, 5, 0, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					Tags: []*Tag{
						{ID: TestTagBuchvorstellungID},
						{ID: TestTagDiskussionID},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataIDLang1,
					StreamID: testdataID,
					Lang:     "de",
					Title:    "Mareice Kaiser",
					Subtitle: "Das Unwohlsein der modernen Mutter",
					Short:    `Mütter sollen heute alles sein: Versorgerin, Businesswoman, Mom I'd like to fuck. Dass darunter ihr Wohlbefinden leidet, ist kein Wunder. Mareice Kaiser, Journalistin und selbst Mutter, stellt dabei immer wieder fest: ...`,
					Long:     `Mütter sollen heute alles sein: Versorgerin, Businesswoman, Mom I'd like to fuck. Dass darunter ihr Wohlbefinden leidet, ist kein Wunder. Mareice Kaiser, Journalistin und selbst Mutter, stellt dabei immer wieder fest: Das Mutterideal ist unerreichbar und voller Widersprüche. Nichts kann man richtig machen und niemandem etwas recht. Mutterschaft berührt dabei, natürlich, jeden Lebensbereich: Denn egal, ob es um Arbeit, Geld, Sex, Körper, Psyche oder Liebe geht – Stereotype, Klischees und gesellschaftlichen Druck gibt es überall, auf Instagram, im Bett und im Büro. In ihrem Buch "Das Unwohlsein der modernen Mutter" (Rowohlt, 2021) zeigt die Autorin, wo Mütter heute stehen: noch immer öfter am Herd als in den Chefetagen. Und, wo sie stehen sollten: Dort, wo sie selbst sich sehen – frei und selbstbestimmt. Bei OUT LOUD liest Mareice Kaiser aus ihrem Buch und spricht mit uns über Frausein, Mutterschaft und Selbstbestimmung.`,
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
