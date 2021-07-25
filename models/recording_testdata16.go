package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Grand Piano: Motschmann Trio
	testdataID := uuid.MustParse("710b445a-51f5-4a9c-8fd4-59956453401c")
	testdataIDLang1 := uuid.MustParse("6c0c2394-bc5c-43b0-ae14-8957118f8231")
	testdataIDSpeaker1 := uuid.MustParse("54e80294-b92d-4dfe-9c67-23e1824b8e4f")
	testdataIDFormat1 := uuid.MustParse("06498c33-8ee8-40a4-b483-6f431fec8112")

	// see stream 11

	/* WARNING unreleased:
	- Public
	- Private
	- Duration
	- Bytes
	*/

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-recording-16",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					CreatedAt: time.Date(2021, 7, 23, 18, 0, 0, 0, loc),
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
							Name:    "Motschmann Trio",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
					Lang:        "de",
					Title:       "Motschmann Trio",
					Subtitle:    "Freitag ab 18:00",
					Short:       `Mit der Premiere der Electric Fields in der Berghain Kantine startete das Trio im Jahr 2016 seine Mission`,
					Long: `**Motschmann Trio**: Mit der Premiere der Electric Fields in der Berghain Kantine startete das Trio im Jahr 2016 seine Mission: Elektronische Musik mit Originalinstrumenten so auf die Bühne zu bringen, dass alle Sounds live generiert werden. Mit einem riesigen Arsenal von analogen Synthesizern, E-Pianos und einem komplexen Multipercussion-Setup reisen Johannes Motschmann, Boris Bolles und David Panzl zwischen den Welten und treten in Clubs wie klassischen Konzertsälen gleichermaßen auf. Electric Fields ist ein Soundtrack, der von leisen Klavierklängen bis zu harten polymetrischen Beats reicht. Ambient- und Dronesounds wechseln sich mit orchestral anmutenden Passagen ab.

Mit einem alten Wurlitzerpiano, das Johannes und Boris im Keller ihres Studentenwohnheims aufgetan hatten, begann die Suche nach immer neuen Instrumenten, die den elektrischen Feldern Jahre später Klang und Gestalt gegeben haben. Im Zentrum stehen die Klänge des Wurlitzerpianos und der CP-70, die mit Bassklängen von Moog Prodigy und MS-20 das harmonische Fundament bilden. Ein gewisser retrospektiver Sound entsteht dadurch, dass fast alle Instrumente aus den 70er und 80er Jahren stammen.

David erweckt Rhythmen zum Leben, die vorab auf Drumcomputern konzipiert wurden und taucht sie in ein neues Licht, während Boris immer wieder mit der Violine einen zerbrechlichen Klang findet, der die Rhythmen und Pattern in einen sphärischen Sound führt. Alles was maschinengesteuert war, liegt nun wieder in den Händen der drei klassisch ausgebildeten Musiker, die mit hoher Präzision Johannes' akribisch ausnotierte Kompositionen so symphonisch klingen lassen, als würde man nicht einem elektroakustischen Trio sondern einem ganzen Orchester zuhören.`,
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
