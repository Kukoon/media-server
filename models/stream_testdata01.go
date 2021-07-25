package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {

	testdataStream := uuid.MustParse("dffe2c0e-3713-4399-8ee2-279becbbb06e")
	testdataStreamLang1 := uuid.MustParse("3a4f9157-65bf-4d15-a82b-1cd9295d07e0")
	testdataStreamSpeaker1 := uuid.MustParse("0030a199-c771-489a-88a7-258f80db2bce")
	testdataStreamSpeaker2 := uuid.MustParse("8bb5af2a-6e66-488b-9eac-6714ce005899")
	testdataStreamSpeaker3 := uuid.MustParse("fa5323fc-5f54-487c-b5cc-173faa4e64f2")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-stream-01",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream,
					ChannelID: TestChannelID1,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 3, 4, 18, 30, 0, 0, loc),
					ListenAt:  time.Date(2021, 3, 4, 18, 30, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataStream.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataStream.String() + "/preview.webp",
					Tags:      []*Tag{{ID: TestTagVortragID}},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataStreamSpeaker1,
							Name:    "Andreas Ehresmann",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataStreamSpeaker2,
							Name:    "Ronald Sperling",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataStreamSpeaker3,
							Name:    "Ines Dirolf",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStreamLang1,
					StreamID: testdataStream,
					Lang:     "de",
					Title:    "„Die mir von der Wehrmacht angebotenen Kriegsgefangenen sind derart entkräftet“",
					Subtitle: "Sowjetische Kriegsgefangene in Bremer Arbeitskommandos 1941-1945",
					Short:    `Sowjetische Kriegsgefangene bildeten eine der größten Opfergruppen des Nationalsozialismus. Die Wehrmacht brachte Millionen sowjetische Soldat\*innen zum Arbeitseinsatz ins Deutsche Reich. Mehr als die Hälfte von ihnen überlebte die Kriegsgefangenschaft in den Kriegsgefangenenlagern wie dem Stalag X B Sandbostel und den Außenkommandos nicht. Auch in Bremen setzten Firmen und Behörden die kriegsgefangenen Rotarmisten zur Arbeit ein, vornehmlich in der Rüstungsindustrie. Im unserem Vortrag wollen wir die ökonomischen und ideologischen Hintergründe und Widersprüche dieser Arbeitseinsätze aufzeigen. ...`,
					Long: `Sowjetische Kriegsgefangene bildeten eine der größten Opfergruppen des Nationalsozialismus. Die Wehrmacht brachte Millionen sowjetische Soldat\*innen zum Arbeitseinsatz ins Deutsche Reich. Mehr als die Hälfte von ihnen überlebte die Kriegsgefangenschaft in den Kriegsgefangenenlagern wie dem Stalag X B Sandbostel und den Außenkommandos nicht. Auch in Bremen setzten Firmen und Behörden die kriegsgefangenen Rotarmisten zur Arbeit ein, vornehmlich in der Rüstungsindustrie. Im unserem Vortrag wollen wir die ökonomischen und ideologischen Hintergründe und Widersprüche dieser Arbeitseinsätze aufzeigen. Anhand einzelner exemplarischer Arbeitskommandos beleuchten wir die Lebens- und Arbeitsbedingungen von sowjetischen Kriegsgefangenen in Bremen. Der Vortrag lädt alle Interessierte zum Austausch über dieses lange verdrängte Thema ein.

Eine Veranstaltung der Gedenkstätte Lager Sandbostel in Kooperation mit dem Kulturzentrum Kukoon.

Bildinfo: Personalkarte des sowjetischen Kriegsgefangenen Wasilij M. Alexejew, der am 15.09.1942 in das Arbeitskommando der Bremer Francke-Werke eingesetzt wurde und am 11.03.1942 an Tuberkulose starb, Quelle [https://obd-memorial.ru/html/info.htm?id=300643349](https://obd-memorial.ru/html/info.htm?id=300643349)`,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Delete(&Stream{
					ID: testdataStream,
				}).Error
			},
		},
	}...)
}
