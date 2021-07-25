package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	testdataStream := uuid.MustParse("57de7dfd-c060-4da1-8f57-f0880c1f2e5e")
	testdataStreamLang1 := uuid.MustParse("d92fbc3b-a027-49f6-849b-7efb425aa5c0")
	testdataStreamSpeaker1 := uuid.MustParse("0d1b38cd-561c-4db4-b4b9-51f74ba3dba4")
	testdataStreamSpeaker2 := uuid.MustParse("1dbf0438-a9c1-4412-b44c-08fe7819902c")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-stream-04",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream,
					ChannelID: TestChannelID1,
					Chat:      true,
					Running:   true,
					StartAt:   time.Date(2021, 4, 22, 0, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 4, 22, 0, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataStream.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataStream.String() + "/preview.webp",
					Tags: []*Tag{
						{ID: TestTagBuchvorstellungID},
						{ID: TestTagDiskussionID},
					},
					Speakers: []*Speaker{
						{
							OwnerID:      TestChannelID1,
							ID:           testdataStreamSpeaker1,
							Name:         "Caro Keller",
							Organisation: "NSU-Watch",
						},
						{
							OwnerID:      TestChannelID1,
							ID:           testdataStreamSpeaker2,
							Name:         "Lee Hielscher",
							Organisation: "Initiative in Gedenken an Nguyễn Ngọc Châu und Đỗ Anh Lân",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStreamLang1,
					StreamID: testdataStream,
					Lang:     "de",
					Title:    "Aufklären und Einmischen",
					Subtitle: "Der NSU-Komplex und der Münchner Prozess - Buchvorstellung mit NSU-Watch",
					Short:    `Im November 2011 kam eine rechtsterroristische Mord- und Anschlagsserie des sogenannten Nationalsozialistischen Untergrunds (NSU) ans Licht, die in ihrer Dimension neu war. In den folgenden Untersuchungen formte sich ein erstes Bild des NSU-Komplexes. ...`,
					Long: `Im November 2011 kam eine rechtsterroristische Mord- und Anschlagsserie des sogenannten Nationalsozialistischen Untergrunds (NSU) ans Licht, die in ihrer Dimension neu war. In den folgenden Untersuchungen formte sich ein erstes Bild des NSU-Komplexes. Dabei wurde deutlich, dass eine noch umfassendere juristische und gesellschaftliche Aufarbeitung anstand. So beschlossen antifaschistische Initiativen und Einzelpersonen, die Arbeit am NSU-Komplex zu verstetigen, und gründeten »NSU-Watch«. Neun Jahre später ist die Aufarbeitung des NSU-Komplexes noch lange nicht abgeschlossen, die Gefahr des rechten Terrors bleibt schrecklich aktuell. NSU-Watch hat den NSU-Prozess beobachtet, jeden Tag protokolliert und der Öffentlichkeit zur Verfügung gestellt. Darüber hinaus haben sich Landesprojekte gegründet, die die parlamentarischen Aufklärungsbemühungen begleiten. Das zentrale Anliegen des Buches von NSU-Watch ist, die rassistischen Strukturen, die den NSU hervorbrachten, ihn wissentlich oder unwissentlich unterstützten und so zehn Morde, drei Sprengstoffanschläge und 15 Raubüberfälle zwischen 1998 und 2011 möglich machten, entlang der Geschehnisse und Akteur*innen des NSU-Prozesses in München aufzuzeigen. Trotz der vielen offen gebliebenen Fragen soll das Buch eine Zwischenbilanz bieten, die antifaschistischer Demokratieförderung zugrunde gelegt werden kann.

## NSU-Watch
Das Autor\*innen-Kollektiv NSU-WATCH besteht aus Mitgliedern der unabhängigen Beobachtungsstelle NSU-Watch – Aufklären & Einmischen, die sich im Jahr 2012 gegründet hat, um die Aufklärungsbemühungen zum NSU-Komplex zu unterstützen und kritisch zu begleiten. NSU-Watch wird von einem Bündnis aus rund einem Dutzend antifaschistischer und antirassistischer Gruppen und Einzelpersonen aus dem ganzen Bundesgebiet getragen, die teilweise seit Jahrzehnten zum Themenkomplex Rechter Terror arbeiten. Kern der Arbeit von NSU-Watch war bzw. ist die Beobachtung des NSU-Prozesses am Oberlandesgericht in München sowie der diversen parlamentarischen Untersuchungsausschüsse im Bundestag und in den Ländern.

Eine Veranstaltung des _Kulturzentrum Kukoon_ in Kooperation mit der _Rosa-Luxemburg-Initiative – Die Rosa-Luxemburg-Stiftung in Bremen_.
				`,
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
