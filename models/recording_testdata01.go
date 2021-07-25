package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	testdataRecording := uuid.MustParse("542685cb-3693-e720-a957-f008f5dae3ee")
	testdataRecordingLang1 := uuid.MustParse("03d33e6a-151f-47d9-be79-a726e0f9a859")
	testdataRecordingSpeaker1 := uuid.MustParse("7998f75e-7252-4ccd-8cfe-06989db28b51")
	testdataRecordingSpeaker2 := uuid.MustParse("c02bb21c-c8dc-4657-b14c-d0625188f463")
	testdataRecordingSpeaker3 := uuid.MustParse("6a3f1c47-9173-479e-bc53-262ea01a3ac1")
	testdataRecordingFormat1 := uuid.MustParse("3a4f9157-65bf-4d15-a82b-1cd9295d07e0")
	testdataRecordingFormat2 := uuid.MustParse("95ac279c-9ec8-4463-9e55-b18f6e6094be")
	testdataRecordingFormat3 := uuid.MustParse("cfcd54de-dc30-4ee9-8877-92515d337af9")
	testdataRecordingFormat4 := uuid.MustParse("e3caa805-ff00-48aa-9410-c939804d5eac")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0020-01-recording-1",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataRecording,
					ChannelID:  TestChannelID1,
					CommonName: "2020-12-polizeigewalt",
					Poster:     "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataRecording.String() + "/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataRecording.String() + "/preview.webp",
					CreatedAt:  time.Date(2020, 12, 10, 20, 0, 0, 0, loc),
					Duration:   time.Hour + 20*time.Minute + 17*time.Second,
					Public:     true,
					Listed:     true,
					Tags:       []*Tag{{ID: TestTagDiskussionID}},
					Speakers: []*Speaker{
						{
							OwnerID:      TestChannelID1,
							ID:           testdataRecordingSpeaker1,
							Name:         "Leila Abdul-Rahman",
							Organisation: "Ruhr Universität Bochum",
						},
						{
							OwnerID:      TestChannelID1,
							ID:           testdataRecordingSpeaker2,
							Name:         "Greta",
							Organisation: "Grün-Weiße Hilfe",
						},
						{
							OwnerID:      TestChannelID1,
							ID:           testdataRecordingSpeaker3,
							Name:         "Mathilda",
							Organisation: "KOP Bremen",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataRecordingLang1,
					RecordingID: testdataRecording,
					Lang:        "de",
					Title:       "Polizeigewalt",
					Subtitle:    "ein deutsches Problem",
					Short:       "Nachdem Mord an George Floyd ist es zu großen Protesten in den Vereinigten Staaten gekommen. Auch in Deutschland sterben schwarze Menschen in Polizeigewahrsam. Ihre Namen sind weitgehend unbekannt: William Tonou-Mbobda, Hussam Fadl, Rooble Warsame, Oury Jalloh, Yaya Diabi, Amed A., Aamir Ageeb, Achidi John, Laya-Alama Condé, Mohamed Idrissi – die Liste ließe sich fortsetzen. ...",
					Long: `Ein deutsches Problem Diskussionsveranstaltung mit Laila Abdul-Rahman, Greta ([Grün-Weiße Hilfe Bremen](https://twitter.com/fanhilfe_bremen?lang=de)) und Mathilda ([Kampagne für Opfer rassistischer Polizeigewalt – KOP Bremen](https://www.facebook.com/KOP-Bremen-Kampagne-f%C3%BCr-Opfer-rassistischer-Polizeigewalt-Bremen-168776953679814/))

Nachdem Mord an George Floyd ist es zu großen Protesten in den Vereinigten Staaten gekommen. Auch in Deutschland sterben schwarze Menschen in Polizeigewahrsam. Ihre Namen sind weitgehend unbekannt: William Tonou-Mbobda, Hussam Fadl, Rooble Warsame, Oury Jalloh, Yaya Diabi, Amed A., Aamir Ageeb, Achidi John, Laya-Alama Condé, Mohamed Idrissi – die Liste ließe sich fortsetzen.
Gemeinsam mit **Laila Abdul-Rahman** vom [Forschungsprojekt Körperverletzung im Amt an der Ruhr-Universität Bochum](https://kviapol.rub.de/index.php/inhalte/zweiter-zwischenbericht), der Grün-Weißen Hilfe und der Kampagne für Opfer rassistischer Polizeigewalt (KOP-Bremen) wollen wir die Themen Polizeigewalt und rassistische Polizeigewalt beleuchten.
Beginnt Polizeigewalt nicht schon bei der sogenannten „Anlasslosen Kontrolle“ oder dem rechtswidrigen Kessel? Warum trifft sie manche Menschen häufiger als andere? Wie geht die Polizei mit Kritik um? Und was unterscheidet die Gewalterfahrungen von Personen mit Migrationshintergrund sowie People of Color von Personen ohne Migrationshintergrund bzw.weißen Personen?

Eine Veranstaltung des Kulturkombinat Offene Neustandt (Kukoon) in Kooperation mit der Rosa-Luxemburg-Initiative – Die Rosa-Luxemburg-Stiftung in Bremen und Partnerschaft für Demokratie Bremen.

## Linksammlung:
Forschungsprojekt Körperverletzung im Amt an der Ruhr-Universität Bochum:
[https://kviapol.rub.de/index.php/inhalte/zweiter-zwischenbericht](https://kviapol.rub.de/index.php/inhalte/zweiter-zwischenbericht)

Grün-Weiße Hilfe e.V.:
[https://twitter.com/fanhilfe_bremen](https://twitter.com/fanhilfe_bremen?lang=de)

Kampagne für Opfer rassistischer Polizeigewalt – KOP Bremen
[kopbremen.noblogs.org](https://kopbremen.noblogs.org/)

### Spenden

Wir bitten um eine kleine Spende an den Verein des Kukoon (*Verein für Bunte Kombinationen e.V.*) per Überweisung an:

**DE72 4306 0967 2063 2646 00**
	
oder per [Paypal](https://www.paypal.com/donate?hosted_button_id=4BQQNN582WLN6) an verein@kukoon.de.
				`,
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&[]*RecordingFormat{
					{
						ID:          testdataRecordingFormat1,
						RecordingID: testdataRecording,
						Lang:        "de",
						Quality:     0,
						IsVideo:     true,
						URL:         "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataRecording.String() + "/video_best.mp4",
						Bytes:       3323919713,
						Resolution:  "1920x1080",
					},
					{
						ID:          testdataRecordingFormat2,
						RecordingID: testdataRecording,
						Lang:        "de",
						Quality:     160,
						IsVideo:     true,
						URL:         "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataRecording.String() + "/video_720.mp4",
						Bytes:       1149359246,
						Resolution:  "1280x720",
					},
					{
						ID:          testdataRecordingFormat3,
						RecordingID: testdataRecording,
						Lang:        "de",
						Quality:     180,
						IsVideo:     true,
						URL:         "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataRecording.String() + "/video_480.mp4",
						Bytes:       654217779,
						Resolution:  "854x480",
					},
					{
						ID:          testdataRecordingFormat4,
						RecordingID: testdataRecording,
						Lang:        "de",
						Quality:     0,
						IsVideo:     false,
						URL:         "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataRecording.String() + "/audio_best.mp4",
						Bytes:       130761076,
						Resolution:  "128kb",
					},
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Delete(&Recording{
					ID: testdataRecording,
				}).Error
			},
		},
	}...)
}
