package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Grand Piano: Pablo Ortega
	testdataID := uuid.MustParse("36ba6bfe-2b40-425d-8cc7-d7de5ec4b67a")
	testdataIDLang1 := uuid.MustParse("c6720db6-9d62-483d-a6ec-5cd1137b4dac")
	testdataIDSpeaker1 := uuid.MustParse("e7e6eb6b-8188-4169-ae4e-8144d559a592")
	testdataIDFormat1 := uuid.MustParse("a7233112-636e-4f81-aa5e-25a0e48e68f6")

	// see stream 09

	/* WARNING unreleased:
	- Public
	- Private
	- Duration
	- Bytes
	*/

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-recording-14",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					CreatedAt: time.Date(2021, 7, 22, 18, 0, 0, 0, loc),
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
							Name:    "Pablo Ortega",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
					Lang:        "de",
					Title:       "Pablo Ortega",
					Subtitle:    "Donnerstag ab 18:00",
					Short:       `Pablo Ortega ist ein spanischer Cellist und Komponist. Er lebt in Bremen, wo er aktuell als multidisziplinarischer Musiker aktiv ist.`,
					Long: `Den zweiten Abend wird **Pablo Ortega** eröffnen.

**Pablo Ortega** ist ein spanischer Cellist und Komponist. Er lebt in Bremen, wo er aktuell als multidisziplinarischer Musiker aktiv ist. In seinen Werken verbindet er Elemente moderner klassischer Musik – wie z.B. intime Cellomelodien – mit elektronischen Beats und organischen, atmosphärischen Klangtexturen mit dem Synthesizer. Damit schafft er eine Mischung von Genres, die von filmischer akustischer Musik bis zu energetischer Electronica reicht.

Seine erste EP »Still Waters Run Deep« erschien im Februar 2020. `,
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
