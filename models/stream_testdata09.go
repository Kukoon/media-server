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

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-stream-09",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 7, 22, 17, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 7, 22, 17, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
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
				if err := tx.Create(&StreamLang{
					ID:       testdataIDLang1,
					StreamID: testdataID,
					Lang:     "de",
					Title:    "Pablo Ortega",
					Subtitle: "Donnerstag ab 18:00",
					Short:    `Pablo Ortega ist ein spanischer Cellist und Komponist. Er lebt in Bremen, wo er aktuell als multidisziplinarischer Musiker aktiv ist.`,
					Long: `Den zweiten Abend wird **Pablo Ortega** eröffnen.

**Pablo Ortega** ist ein spanischer Cellist und Komponist. Er lebt in Bremen, wo er aktuell als multidisziplinarischer Musiker aktiv ist. In seinen Werken verbindet er Elemente moderner klassischer Musik – wie z.B. intime Cellomelodien – mit elektronischen Beats und organischen, atmosphärischen Klangtexturen mit dem Synthesizer. Damit schafft er eine Mischung von Genres, die von filmischer akustischer Musik bis zu energetischer Electronica reicht.

Seine erste EP »Still Waters Run Deep« erschien im Februar 2020. `,
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
