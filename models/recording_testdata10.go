package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// see stream 05
	testdataID := uuid.MustParse("4fb029d6-063a-4302-9ae8-4c1c6a1542a5")
	testdataIDLang1 := uuid.MustParse("d5262bb7-378b-456f-9e91-34f63b174c48")
	testdataIDFormat1 := uuid.MustParse("5473a466-3a71-4be3-8436-a34f92c5ecc6")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0020-01-recording-10",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataID,
					ChannelID:  TestChannelID1,
					EventID:    &TestEventID1,
					CommonName: "2021-05_out_loud-mareice_kaiser-modernen_mutter",
					Poster:     "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					CreatedAt:  time.Date(2021, 5, 5, 19, 0, 0, 0, loc),
					Duration:   time.Hour + 28*time.Minute + 26*time.Second,
					Public:     true,
					Listed:     true,
					Tags: []*Tag{
						{ID: TestTagBuchvorstellungID},
						{ID: TestTagDiskussionID},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
					Lang:        "de",
					Title:       "Mareice Kaiser",
					Subtitle:    "Das Unwohlsein der modernen Mutter",
					Short:       `Mütter sollen heute alles sein: Versorgerin, Businesswoman, Mom I'd like to fuck. Dass darunter ihr Wohlbefinden leidet, ist kein Wunder. Mareice Kaiser, Journalistin und selbst Mutter, stellt dabei immer wieder fest: ...`,
					Long:        `Mütter sollen heute alles sein: Versorgerin, Businesswoman, Mom I'd like to fuck. Dass darunter ihr Wohlbefinden leidet, ist kein Wunder. Mareice Kaiser, Journalistin und selbst Mutter, stellt dabei immer wieder fest: Das Mutterideal ist unerreichbar und voller Widersprüche. Nichts kann man richtig machen und niemandem etwas recht. Mutterschaft berührt dabei, natürlich, jeden Lebensbereich: Denn egal, ob es um Arbeit, Geld, Sex, Körper, Psyche oder Liebe geht – Stereotype, Klischees und gesellschaftlichen Druck gibt es überall, auf Instagram, im Bett und im Büro. In ihrem Buch "Das Unwohlsein der modernen Mutter" (Rowohlt, 2021) zeigt die Autorin, wo Mütter heute stehen: noch immer öfter am Herd als in den Chefetagen. Und, wo sie stehen sollten: Dort, wo sie selbst sich sehen – frei und selbstbestimmt. Bei OUT LOUD liest Mareice Kaiser aus ihrem Buch und spricht mit uns über Frausein, Mutterschaft und Selbstbestimmung.`,
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
						Bytes:       2443666130,
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
