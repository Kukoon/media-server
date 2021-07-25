package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Grand Piano: Kat Kit
	testdataID := uuid.MustParse("88b00ee3-e528-423c-91d4-21b9cb724c04")
	testdataIDLang1 := uuid.MustParse("fd9da8c0-3ac8-49ed-b915-b6a76c2f5ccb")
	testdataIDSpeaker1 := uuid.MustParse("36ff226f-42d6-4d48-9cce-4f504ba2d0b5")
	testdataIDFormat1 := uuid.MustParse("f44003a2-f8f5-4f75-bf6d-112ab1b337f4")

	// see stream 16

	/* WARNING unreleased:
	- Public
	- Private
	- Duration
	- Bytes
	*/

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-recording-21",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					CreatedAt: time.Date(2021, 7, 25, 16, 0, 0, 0, loc),
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
							Name:    "Kat Kit",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
					Lang:        "de",
					Title:       "Kat Kit",
					Subtitle:    "Sonntag ab 16:00",
					Short:       `**Kat Kit** ist eine Singer-Songwriterin, die keine sein will.`,
					Long: `Den Auftakt des letzten Tages macht **Kat Kit**:
**Kat Kit** ist eine Singer-Songwriterin, die keine sein will.
Sie hat durch Peer Pressure mit der Musik angefangen - im Kinderchor gab es nämlich Eis für alle, wenn genug Kids beitreten. Dann ist sie einfach dabei geblieben, schließlich gab es sowieso nicht viel mehr zu tun in dem 500-Seelen Dorf nahe Kaiserslautern. Bis heute hat sich nichts besseres gefunden, Kat studiert mittlerweile Popkomposition an der Folkwang Universität der Künste Essen. Ihr Sound befindet sich irgendwo zwischen metaphorischen Traumlandschaften, Regina Spektor und Studio Ghibli Ästhetik. Wie Jazzpop, nur mit eingängigeren Melodien, dazu pittoreske Texte mit einer ordentlichen Portion Dramatik und einem Hauch Minimalismus. **Kat Kit** ist ein Balanceakt aus Intimität und Ironie. Ein Raum, um Zwängen zu entfliehen und sie gleichermaßen lyrisch zu entkräften. Die Bandbreite reicht von Whatsapp-Sprachnachrichten Performances bis zu majestätischer Musik über verbotene Früchte. Hauptsache, kein Mainstream.`,
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
