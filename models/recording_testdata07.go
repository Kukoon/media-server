package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// see stream 01
	testdataID := uuid.MustParse("728edaf7-9ad9-f972-4d09-ba5940cd43f9")
	testdataIDLang1 := uuid.MustParse("acdf7eb1-0cb9-4900-a918-a411f9afc38d")
	testdataIDSpeaker1 := uuid.MustParse("0030a199-c771-489a-88a7-258f80db2bce")
	testdataIDSpeaker2 := uuid.MustParse("8bb5af2a-6e66-488b-9eac-6714ce005899")
	testdataIDSpeaker3 := uuid.MustParse("fa5323fc-5f54-487c-b5cc-173faa4e64f2")
	testdataIDFormat1 := uuid.MustParse("4069206c-e6e5-4320-ab12-74af566791e3")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0020-01-recording-7",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataID,
					ChannelID:  TestChannelID1,
					CommonName: "2021-03-kriegsgefanngende_in_bremen",
					Poster:     "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					CreatedAt:  time.Date(2021, 3, 4, 19, 0, 0, 0, loc),
					Duration:   time.Hour + 4*time.Minute + 25*time.Second,
					Public:     true,
					Listed:     true,
					Tags: []*Tag{
						{ID: TestTagVortragID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker1,
							Name:    "Andreas Ehresmann",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker2,
							Name:    "Ronald Sperling",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker3,
							Name:    "Ines Dirolf",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
					Lang:        "de",
					Title:       "„Die mir von der Wehrmacht angebotenen Kriegsgefangenen sind derart entkräftet“",
					Subtitle:    "Sowjetische Kriegsgefangene in Bremer Arbeitskommandos 1941-1945",
					Short:       `Sowjetische Kriegsgefangene bildeten eine der größten Opfergruppen des Nationalsozialismus. Die Wehrmacht brachte Millionen sowjetische Soldat\*innen zum Arbeitseinsatz ins Deutsche Reich. ...`,
					Long: `
Sowjetische Kriegsgefangene bildeten eine der größten Opfergruppen des Nationalsozialismus. Die Wehrmacht brachte Millionen sowjetische Soldat\*innen zum Arbeitseinsatz ins Deutsche Reich. Mehr als die Hälfte von ihnen überlebte die Kriegsgefangenschaft in den Kriegsgefangenenlagern wie dem Stalag X B Sandbostel und den Außenkommandos nicht. Auch in Bremen setzten Firmen und Behörden die kriegsgefangenen Rotarmisten zur Arbeit ein, vornehmlich in der Rüstungsindustrie. Im unserem Vortrag wollen wir die ökonomischen und ideologischen Hintergründe und Widersprüche dieser Arbeitseinsätze aufzeigen. Anhand einzelner exemplarischer Arbeitskommandos beleuchten wir die Lebens- und Arbeitsbedingungen von sowjetischen Kriegsgefangenen in Bremen. Der Vortrag lädt alle Interessierte zum Austausch über dieses lange verdrängte Thema ein.

Online-Vortrag mit Andreas Ehresmann, Ronald Sperling und Ines Dirolf.

Eine Veranstaltung der Gedenkstätte Lager Sandbostel in Kooperation mit dem Kulturzentrum Kukoon.

Bildinfo: Personalkarte des sowjetischen Kriegsgefangenen Wasilij M. Alexejew, der am 15.09.1942 in das Arbeitskommando der Bremer Francke-Werke eingesetzt wurde und am 11.03.1942 an Tuberkulose starb, Quelle [https://obd-memorial.ru/html/info.htm?id=300643349](https://obd-memorial.ru/html/info.htm?id=300643349)
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
						Bytes:       958856106,
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
