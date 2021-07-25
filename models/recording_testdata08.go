package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// see stream 03
	testdataID := uuid.MustParse("06e3a71e-581d-4735-9647-3e4a49b5caa8")
	testdataIDLang1 := uuid.MustParse("cff00fcd-5408-4cb4-8ac7-2d42b45fbd68")
	testdataIDSpeaker1 := uuid.MustParse("c03aa102-551e-4b3a-b670-5b6c7ac13faa")
	testdataIDFormat1 := uuid.MustParse("b98078df-b430-4a19-971c-84d324fd9b14")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0020-01-recording-8",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataID,
					ChannelID:  TestChannelID1,
					CommonName: "2021-03-kriegsgefanngende_in_bremen",
					Poster:     "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					CreatedAt:  time.Date(2021, 4, 15, 19, 0, 0, 0, loc),
					Duration:   time.Hour + 8*time.Minute,
					Public:     true,
					Listed:     true,
					Tags: []*Tag{
						{ID: TestTagBuchvorstellungID},
						{ID: TestTagDiskussionID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker1,
							Name:    "Andreas Speit",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
					Lang:        "de",
					Title:       "Rechte Egoshooter",
					Subtitle:    "Von der virtuellen Hetze zum Livestream-Attentat",
					Short:       `Weltweit gibt es rechtsterroristische Attentate eines neuen Typs. In Halle (Saale) verhinderte nur eine verschlossene Holztür der Synagoge ein größeres Massaker. ...`,
					Long: `Weltweit gibt es rechtsterroristische Attentate eines neuen Typs. In Halle (Saale) verhinderte nur eine verschlossene Holztür der Synagoge ein größeres Massaker. Am 9. Oktober 2019 wollte dort ein Rechtsextremist die versammelten Juden hinrichten. Mit selbstgebauten Waffen schoss er auf die Tür und warf eigens hergestellte Sprengsätze. Online konnten Gleichgesinnte zusehen, wie er zwei Menschen ermordete: Seine Tat verbreitete er per Videokamera auf einem Portal für Computerspiel-Videos. Er ahmte damit andere »Egoshooter« nach – wie einen Rechtsextremisten, der in Neuseeland wenige Monate zuvor die Tötung von 51 Menschen live im Internet übertragen hatte. Was treibt Menschen vom Bildschirm zur realen Gewalt auf der Straße? Die Beiträge des Buches gehen den Spuren der Attentäter nach und zeigen die speziellen Radikalisierungsmechanismen im Netz auf. Sie erklären die Hintergründe und Motive dieser Männer, die in ihren rechten Online-Gemeinden Antisemitismus, Rassismus und Antifeminismus verbreiten. Das Buch gibt Einblicke in eine Welt, die vielen unbekannt ist.

Eine Veranstaltung des _Kulturzentrum Kukoon_ in Kooperation mit der _Rosa-Luxemburg-Initiative – Die Rosa-Luxemburg-Stiftung in Bremen_.
				`,
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
						Bytes:       995280142,
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
