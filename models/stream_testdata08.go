package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Grand Piano: Lucia Cadotsch
	testdataStream := uuid.MustParse("7ff58740-8c3a-4e09-8fc1-1eeb39c2a9d4")
	testdataStreamLang1 := uuid.MustParse("b7be68f2-a109-4e28-8744-bfc6c8f03f9f")
	testdataStreamSpeaker1 := uuid.MustParse("dfa0ff16-8cb0-46e2-a56a-e44dcde1868e")
	testdataStreamSpeaker2 := uuid.MustParse("bd72ee71-6e0f-4ba6-9d3d-eb6b7ba589e3")
	testdataStreamSpeaker3 := testdataSpeakerClaraVetter

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-stream-08",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 7, 21, 18, 15, 0, 0, loc),
					ListenAt:  time.Date(2021, 7, 21, 18, 15, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataStream.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataStream.String() + "/preview.webp",
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataStreamSpeaker1,
							Name:    "Lucia Cadotsch",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataStreamSpeaker2,
							Name:    "Ronny Graupe",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataStreamSpeaker3,
							Name:    "Clara Vetter",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStreamLang1,
					StreamID: testdataStream,
					Lang:     "de",
					Title:    "Lucia Cadotsch",
					Subtitle: "Mittwoch ab 20:30",
					Short: `Hierzu ziteren wir den Deutschlandfunk:
__Sie hat dem Jazzgesang neuen Atem eingehaucht und experimentiert ständig weiter.__ (deutschlandfunk.de)`,
					Long: `Am ersten Abend wird uns **Lucia Cadotsch** begleitet von **Ronny Graupe** (Gitarre) in die Nacht begleiten. Dabei werden sie von **Clara Vetter** unterstützt.

Hierzu ziteren wir den Deutschlandfunk:
__Sie hat dem Jazzgesang neuen Atem eingehaucht und experimentiert ständig weiter.__ (deutschlandfunk.de)

1984 wurde **Lucia Cadotsch** in Zürich geboren. Mit 14 entdeckte sie ihre Liebe zum Jazz: Die Plattensammlung ihres Vaters begeisterte sie, vor allem Aufnahmen mit Miles Davis, John Coltrane, Nina Simone und Billie Holiday.
**Lucia Cadotsch** bekam Klavier- und Gesangsunterricht und ging mit 18 Jahren an die Universität der Künste Berlin, um Jazzgesang zu studieren. Sie begründete diverse Ensembles, zum Beispiel das Popquartett Schneeweiss + Rosenrot, mit dem sie 2012 den Neuen Deutschen Jazzpreis gewann.
2016 gelang **Lucia Cadotsch** der internationale Durchbruch mit dem Album »Speak Low«. Ein Jahr später erhielt sie den ECHO Jazz als Sängerin des Jahres. Heute lebt sie in Berlin. `,
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
