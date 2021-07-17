package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	loc = time.FixedZone("UTC+2", +2*60*60)

	testdataStream1      = uuid.MustParse("dffe2c0e-3713-4399-8ee2-279becbbb06e")
	testdataStream1Lang1 = uuid.MustParse("3a4f9157-65bf-4d15-a82b-1cd9295d07e0")

	testdataStream2 = uuid.MustParse("0801a547-59f1-4a63-946f-2ab03f62e6ee")

	testdataStream3         = uuid.MustParse("06e3a71e-581d-4735-9647-3e4a49b5caa8")
	testdataStream3Lang1    = uuid.MustParse("cff00fcd-5408-4cb4-8ac7-2d42b45fbd68")
	testdataStream3Speaker1 = uuid.MustParse("c03aa102-551e-4b3a-b670-5b6c7ac13faa")

	testdataStream4         = uuid.MustParse("57de7dfd-c060-4da1-8f57-f0880c1f2e5e")
	testdataStream4Lang1    = uuid.MustParse("d92fbc3b-a027-49f6-849b-7efb425aa5c0")
	testdataStream4Speaker1 = uuid.MustParse("0d1b38cd-561c-4db4-b4b9-51f74ba3dba4")
	testdataStream4Speaker2 = uuid.MustParse("1dbf0438-a9c1-4412-b44c-08fe7819902c")

	testdataStream5      = uuid.MustParse("4fb029d6-063a-4302-9ae8-4c1c6a1542a5")
	testdataStream5Lang1 = uuid.MustParse("d5262bb7-378b-456f-9e91-34f63b174c48")

	testdataStream6      = uuid.MustParse("1742d9b6-c9c6-45fb-a3a3-4a3e7fac2987")
	testdataStream6Lang1 = uuid.MustParse("0b7136a6-4c51-49ac-99e9-27ef833169f6")

	testdataRecording1         = uuid.MustParse("542685cb-3693-e720-a957-f008f5dae3ee")
	testdataRecording1Lang1    = uuid.MustParse("03d33e6a-151f-47d9-be79-a726e0f9a859")
	testdataRecording1Format1  = uuid.MustParse("3a4f9157-65bf-4d15-a82b-1cd9295d07e0")
	testdataRecording1Format2  = uuid.MustParse("95ac279c-9ec8-4463-9e55-b18f6e6094be")
	testdataRecording1Format3  = uuid.MustParse("cfcd54de-dc30-4ee9-8877-92515d337af9")
	testdataRecording1Format4  = uuid.MustParse("e3caa805-ff00-48aa-9410-c939804d5eac")
	testdataRecording1Speaker1 = uuid.MustParse("7998f75e-7252-4ccd-8cfe-06989db28b51")
	testdataRecording1Speaker2 = uuid.MustParse("c02bb21c-c8dc-4657-b14c-d0625188f463")
	testdataRecording1Speaker3 = uuid.MustParse("6a3f1c47-9173-479e-bc53-262ea01a3ac1")

	testdataRecording2         = uuid.MustParse("45da89a7-e5e0-5104-b937-6d4c2d4b6d00")
	testdataRecording2Lang1    = uuid.MustParse("20eb49e4-46a6-4b30-87eb-72e4dbe77b09")
	testdataRecording2Format1  = uuid.MustParse("09858a43-0532-4ad8-8694-92ed71372ff4")
	testdataRecording2Format2  = uuid.MustParse("6b4cf4b0-db39-47a3-9bc9-4ff6ef539a88")
	testdataRecording2Format3  = uuid.MustParse("fc129d9b-b774-4b47-a7f3-a8cf046a5573")
	testdataRecording2Format4  = uuid.MustParse("918c0386-03ad-4d34-b930-d6ce6a7632eb")
	testdataRecording2Speaker1 = uuid.MustParse("5dfddf98-2b11-48c7-ab5e-8cc6a2db0dcd")
	testdataRecording2Speaker2 = uuid.MustParse("a85fffb1-aa9f-4bd0-81e7-0058891827ef")

	testdataRecording3         = uuid.MustParse("edb1cfbb-3476-d639-b3f5-795fabf4ef4d")
	testdataRecording3Lang1    = uuid.MustParse("d193bc41-99f6-46d8-870e-72a860520223")
	testdataRecording3Format1  = uuid.MustParse("6b1b95f2-d92d-4da7-b56c-1ba86ff22dcd")
	testdataRecording3Speaker1 = uuid.MustParse("a098c2f5-aa63-4c54-87b1-46ddda1cde16")

	testdataRecording4         = uuid.MustParse("13a70ec7-6e74-5114-021f-4d7910752df1")
	testdataRecording4Lang1    = uuid.MustParse("345a3743-42dd-4ee9-97ae-c3785bd4235e")
	testdataRecording4Format1  = uuid.MustParse("fae90633-a15d-4dfb-b017-dc8561df95c3")
	testdataRecording4Speaker1 = uuid.MustParse("f597b912-97b1-4e16-b431-054692a5d049")

	testdataRecording5        = uuid.MustParse("27efbfff-d66c-c935-b308-9b1ee2bf78c8")
	testdataRecording5Lang1   = uuid.MustParse("4f5ad673-2496-429a-a74f-0b48acdb807b")
	testdataRecording5Format1 = uuid.MustParse("357af110-9481-4d0e-9fea-f61b30ee26f4")

	testdataRecording6         = uuid.MustParse("81b262e9-e010-1fa2-84a5-d8cee1a94835")
	testdataRecording6Lang1    = uuid.MustParse("0ce4b366-9238-4aa4-a6d6-94227c1b0681")
	testdataRecording6Format1  = uuid.MustParse("449e3361-f2e2-44ee-a5d7-3c013cfe1fdc")
	testdataRecording6Speaker1 = uuid.MustParse("d8ba2b91-78f7-4bcd-9dc4-5af1d3c904a9")
	testdataRecording6Speaker2 = uuid.MustParse("62d9ce45-1465-40f8-bf99-22607e7be91d")

	testdataRecording7         = uuid.MustParse("728edaf7-9ad9-f972-4d09-ba5940cd43f9")
	testdataRecording7Lang1    = uuid.MustParse("acdf7eb1-0cb9-4900-a918-a411f9afc38d")
	testdataRecording7Format1  = uuid.MustParse("4069206c-e6e5-4320-ab12-74af566791e3")
	testdataRecording7Speaker1 = uuid.MustParse("0030a199-c771-489a-88a7-258f80db2bce")
	testdataRecording7Speaker2 = uuid.MustParse("8bb5af2a-6e66-488b-9eac-6714ce005899")
	testdataRecording7Speaker3 = uuid.MustParse("fa5323fc-5f54-487c-b5cc-173faa4e64f2")

	testdataRecording8         = testdataStream3
	testdataRecording8Lang1    = testdataStream3Lang1
	testdataRecording8Speaker1 = testdataStream3Speaker1
	testdataRecording8Format1  = uuid.MustParse("b98078df-b430-4a19-971c-84d324fd9b14")

	testdataRecording9         = testdataStream4
	testdataRecording9Lang1    = testdataStream4Lang1
	testdataRecording9Speaker1 = testdataStream4Speaker1
	testdataRecording9Speaker2 = testdataStream4Speaker2
	testdataRecording9Format1  = uuid.MustParse("e0250ff4-ad36-47d2-a58a-5ba857f50ab4")

	testdataRecording10        = testdataStream5
	testdataRecording10Lang1   = testdataStream5Lang1
	testdataRecording10Format1 = uuid.MustParse("5473a466-3a71-4be3-8436-a34f92c5ecc6")
)

func init() {

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0020-01-recording-1",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataRecording1,
					ChannelID:  TestChannelID1,
					CommonName: "2020-12-polizeigewalt",
					Poster:     "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee/preview.webp",
					CreatedAt:  time.Date(2020, 12, 10, 20, 0, 0, 0, loc),
					Duration:   time.Hour + 20*time.Minute + 17*time.Second,
					Public:     true,
					Listed:     true,
					Tags:       []*Tag{{ID: TestTagDiskussionID}},
					Speakers: []*Speaker{
						{
							OwnerID:      TestChannelID1,
							ID:           testdataRecording1Speaker1,
							Name:         "Leila Abdul-Rahman",
							Organisation: "Ruhr Universität Bochum",
						},
						{
							OwnerID:      TestChannelID1,
							ID:           testdataRecording1Speaker2,
							Name:         "Greta",
							Organisation: "Grün-Weiße Hilfe",
						},
						{
							OwnerID:      TestChannelID1,
							ID:           testdataRecording1Speaker3,
							Name:         "Mathilda",
							Organisation: "KOP Bremen",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataRecording1Lang1,
					RecordingID: testdataRecording1,
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
				if err := tx.Create(&RecordingFormat{
					ID:          testdataRecording1Format1,
					RecordingID: testdataRecording1,
					Lang:        "de",
					Quality:     0,
					IsVideo:     true,
					URL:         "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee/video_best.mp4",
					Bytes:       3323919713,
					Resolution:  "1920x1080",
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingFormat{
					ID:          testdataRecording1Format2,
					RecordingID: testdataRecording1,
					Lang:        "de",
					Quality:     160,
					IsVideo:     true,
					URL:         "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee/video_720.mp4",
					Bytes:       1149359246,
					Resolution:  "1280x720",
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingFormat{
					ID:          testdataRecording1Format3,
					RecordingID: testdataRecording1,
					Lang:        "de",
					Quality:     180,
					IsVideo:     true,
					URL:         "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee/video_480.mp4",
					Bytes:       654217779,
					Resolution:  "854x480",
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingFormat{
					ID:          testdataRecording1Format4,
					RecordingID: testdataRecording1,
					Lang:        "de",
					Quality:     0,
					IsVideo:     false,
					URL:         "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee/audio_best.mp3",
					Bytes:       130761076,
					Resolution:  "128kb",
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&RecordingFormat{
					ID: testdataRecording1Format4,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&RecordingFormat{
					ID: testdataRecording1Format3,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&RecordingFormat{
					ID: testdataRecording1Format2,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&RecordingFormat{
					ID: testdataRecording1Format1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&RecordingLang{
					ID: testdataRecording1Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Recording{
					ID: testdataRecording1,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0020-01-recording-2",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataRecording2,
					ChannelID:  TestChannelID1,
					CommonName: "2021-01-faschistische_jahrhundert",
					Poster:     "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00/preview.webp",
					CreatedAt:  time.Date(2021, 01, 29, 20, 0, 0, 0, loc),
					Duration:   time.Hour + 43*time.Minute + 10*time.Second,
					Public:     true,
					Listed:     true,
					Tags: []*Tag{
						{ID: TestTagBuchvorstellungID},
						{ID: TestTagDiskussionID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataRecording2Speaker1,
							Name:    "Volkmar Wölk",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataRecording2Speaker2,
							Name:    "Felix Schilk",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataRecording2Lang1,
					RecordingID: testdataRecording2,
					Lang:        "de",
					Title:       "Das faschistische Jahrhundert",
					Subtitle:    "Neurechte Diskurse zu Abendland, Identität, Europa, Neoliberalismus",
					Short:       "„Wer das Wesen des Faschismus erkennen will, muss zurück zu dessen Wurzeln“ ­– so der Faschismusforscher Zeev Sternhell. Gleiches gilt für die europäische Neue Rechte. ...",
					Long: `## Buchvorstellung: Das faschistische Jahrhundert

„Wer das Wesen des Faschismus erkennen will, muss zurück zu dessen Wurzeln“ – so der Faschismusforscher Zeev Sternhell. Gleiches gilt für die europäische Neue Rechte. Aus welchen Gründen ist sie wann und wo mit welchen Zielen entstanden? Wo liegen ihre geistigen Wurzeln? Betreiben die „neuen“ Rechten tatsächlich eine ideologische Erneuerung der extremen Rechten, oder handelt es sich lediglich um alten Wein in neuen Schläuchen? **Volkmar Wölk** und **Felix Schilk**, Mitautoren des Bandes _Das faschistische Jahrhundert_, veranschaulichen die Denkwege der Neuen Rechten anhand von zwei zentralen Themenfelder: Einerseits ihre Europakonzeptionen und andererseits ihre wirtschafts- und sozialpolitischen Vorstellungen.

Ein kurzer Beitrag zu dem Buch [hier](https://www.deutschlandfunkkultur.de/friedrich-burschel-das-faschistische-jahrhundert-schon.1270.de.html?dram:article_id=485508).

Eine Veranstaltung des _Kulturzentrum Kukoon_ in Kooperation mit der _Rosa-Luxemburg-Initiative – Die Rosa-Luxemburg-Stiftung in Bremen_.
				`,
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingFormat{
					ID:          testdataRecording2Format1,
					RecordingID: testdataRecording2,
					Lang:        "de",
					Quality:     0,
					IsVideo:     true,
					URL:         "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00/video_best.mp4",
					Bytes:       2878429977,
					Resolution:  "1920x1080",
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingFormat{
					ID:          testdataRecording2Format2,
					RecordingID: testdataRecording2,
					Lang:        "de",
					Quality:     160,
					IsVideo:     true,
					URL:         "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00/video_720.mp4",
					Bytes:       1149359246,
					Resolution:  "1280x720",
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingFormat{
					ID:          testdataRecording2Format3,
					RecordingID: testdataRecording2,
					Lang:        "de",
					Quality:     180,
					IsVideo:     true,
					URL:         "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00/video_480.mp4",
					Bytes:       480268045,
					Resolution:  "854x480",
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingFormat{
					ID:          testdataRecording2Format4,
					RecordingID: testdataRecording2,
					Lang:        "de",
					Quality:     0,
					IsVideo:     false,
					URL:         "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00/audio_best.mp3",
					Bytes:       115967040,
					Resolution:  "149kb",
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&RecordingFormat{
					ID: testdataRecording2Format4,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&RecordingFormat{
					ID: testdataRecording2Format3,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&RecordingFormat{
					ID: testdataRecording2Format2,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&RecordingFormat{
					ID: testdataRecording2Format1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&RecordingLang{
					ID: testdataRecording2Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Recording{
					ID: testdataRecording2,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0020-01-recording-3",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataRecording3,
					ChannelID:  TestChannelID1,
					CommonName: "2021-02-pushbacks",
					Poster:     "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/edb1cfbb-3476-d639-b3f5-795fabf4ef4d/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/edb1cfbb-3476-d639-b3f5-795fabf4ef4d/preview.webp",
					CreatedAt:  time.Date(2021, 2, 5, 20, 0, 0, 0, loc),
					Duration:   time.Hour + 43*time.Minute + 21*time.Second,
					Public:     true,
					Listed:     true,
					Tags:       []*Tag{{ID: TestTagVortragID}},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataRecording3Speaker1,
							Name:    "Bernd Kasparek",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataRecording3Lang1,
					RecordingID: testdataRecording3,
					Lang:        "de",
					Title:       "Pushbacks, Internierung, Zugangshürden",
					Subtitle:    "Zum Stand des europäischen Migrations- und Grenzregimes",
					Short:       "Nach dem katastrophalen Brand des Flüchtlingslagers Moria auf Lesbos hatte die Europäische Kommission erneut einen Neustart in der europäischen Migrations- und Asylpolitik versucht. ...",
					Long: `Nach dem katastrophalen Brand des Flüchtlingslagers Moria auf Lesbos hatte die Europäische Kommission erneut einen Neustart in der europäischen Migrations- und Asylpolitik versucht. Der Neue Pakt für Asyl und Migration spitzt jedoch vor allem die Entwicklungen zu, die das europäische Asylsystem seit dem Sommer der Migration 2015 kennzeichnen: Zugangshürden zum Asylsystem und massive Internierung. Der Pakt adressiert jedoch nicht die bestehenden Rechtsverletzungen, die das europäische Grenzregime kennzeichnen: illegale Pushbacks sind an vielen Abschnitten der europäischen Außengrenze zur Normalität geworden, und scheinen von der europäischen Grenzschutzagentur Frontex stillschweigend geduldet, oder sogar aktiv unterstützt zu werden. Dies ist vor allem deswegen Besorgnis erregend, da die Agentur bis 2027 über ein eigenes Grenzschutzkorps mit 10.000 Grenzschützern verfügen soll. In der Veranstaltung werden all diese Entwicklungen des letzten Jahres skizziert werden, wobei eine Diskussion um Alternativen und Handlungsmöglichkeiten für ein solidarisches Europa von unten nicht zu kurz kommen sollen.

**Bernd Kasparek** ist seit vielen Jahren Aktivist und Grenzregimeforscher. Sein Buch „Europa als Grenze. Eine Ethnographie der Grenzschutz-Agentur Frontex“ wird im Mai bei transcript erscheinen. Er ist Mitglied des Netzwerks für Kritische Migrations- und Grenzregimeforschung und der Forschungsassoziation [bordermonitoring.eu](http://bordermonitoring.eu/) e.V.

Eine Veranstaltung des Kulturzentrum Kukoon in Kooperation mit der Rosa-Luxemburg-Initiative – Die Rosa-Luxemburg-Stiftung in Bremen.
				`,
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingFormat{
					ID:          testdataRecording3Format1,
					RecordingID: testdataRecording3,
					Lang:        "de",
					Quality:     0,
					IsVideo:     true,
					URL:         "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/edb1cfbb-3476-d639-b3f5-795fabf4ef4d/video_best.mp4",
					Bytes:       1092701356,
					Resolution:  "1920x1080",
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&RecordingFormat{
					ID: testdataRecording3Format1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&RecordingLang{
					ID: testdataRecording3Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Recording{
					ID: testdataRecording3,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0020-01-recording-4",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataRecording4,
					ChannelID:  TestChannelID1,
					CommonName: "2021-02-der_berg_der_nackten_wahrheiten",
					Poster:     "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/13a70ec7-6e74-5114-021f-4d7910752df1/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/13a70ec7-6e74-5114-021f-4d7910752df1/preview.webp",
					CreatedAt:  time.Date(2021, 2, 11, 19, 0, 0, 0, loc),
					Duration:   29*time.Minute + 56*time.Second,
					Public:     true,
					Listed:     true,
					Tags:       []*Tag{{ID: TestTagBuchvorstellungID}},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataRecording4Speaker1,
							Name:    "Jan Backmann",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataRecording4Lang1,
					RecordingID: testdataRecording4,
					Lang:        "de",
					Title:       "Der Berg der nackten Wahrheiten",
					Subtitle:    "Die Geschichte des legendären Monte Verità aus der Sicht einer Ziege erzählt",
					Short:       `Gusto verbringt sein Leben auf dem Monte Verità im Tessin. Doch nicht alles verläuft so sorgenlos, wie der Hippie-Vorläufer es sich in seiner Traumwelt vorstellt hatte: Das Geld der vegetarisch-kommunistischen FKK-Gemeinschaft wird langsam knapp und als Gusto auch noch eine Ziege aus dem Dorfe bei sich aufnimmt, wächst die Wut der Bewohner\*innen von Ascona auf die Aussteiger\*innen.  ...`,
					Long: `
Gusto verbringt sein Leben auf dem Monte Verità im Tessin. Doch nicht alles verläuft so sorgenlos, wie der Hippie-Vorläufer es sich in seiner Traumwelt vorstellt hatte: Das Geld der vegetarisch-kommunistischen FKK-Gemeinschaft wird langsam knapp und als Gusto auch noch eine Ziege aus dem Dorfe bei sich aufnimmt, wächst die Wut der Bewohner\*innen von Ascona auf die Aussteiger\*innen. Nichtsdestotrotz schmiedet Gusto einen irrwitzigen Plan, wie er seine geliebte Ziege weiterhin bei sich behalten kann. Nach seinem ersten, sehr erfolgreichen Comic *Mühsam, Anarchist in Anführungsstrichen*, veröffentlicht der Autor nun eine Erzählung, die zehn Jahre früher spielt, vor dem Hintergrund der Aktivitäten auf dem Monte Verità, dem Treffpunkt der ersten Aussteiger\*innen im 20. Jahrhundert. Auch dieses Mal zapft Bachmann historische Quellen an, um daraus eine pointierte und bissige politische Komödie zu machen. Im Mittelpunkt steht nun allerdings eine Ziege, die Ziege der Vegetarier\*innen. Eine Leseprobe findet sich [hier](https://www.editionmoderne.ch/buch/der-berg-der-nackten-wahrheiten/).

**Jan Bachmann**
Geboren 1986 in Basel, hat an der Deutschen Film- und Fernsehakademie in Berlin studiert. 2013 bis 2015 war er Mitglied in einem brandenburgischen FKK-Verein. Sein erster Comic *Mühsam, Anarchist in Anführungsstrichen* ist 2018 bei der Edition Moderne erschienen und wurde unter anderem für den Max und Moritz-Preis nominiert. Aktuell arbeitet er an einem Buch zum Exil von Kaiser Wilhelm II in Holland.
				`,
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingFormat{
					ID:          testdataRecording4Format1,
					RecordingID: testdataRecording4,
					Lang:        "de",
					Quality:     0,
					IsVideo:     true,
					URL:         "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/13a70ec7-6e74-5114-021f-4d7910752df1/video_best.mp4",
					Bytes:       2020856776,
					Resolution:  "1920x1080",
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&RecordingFormat{
					ID: testdataRecording4Format1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&RecordingLang{
					ID: testdataRecording4Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Recording{
					ID: testdataRecording4,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0020-01-recording-5",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataRecording5,
					ChannelID:  TestChannelID1,
					CommonName: "2021-02-geschichte_wird_gemacht",
					Poster:     "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/27efbfff-d66c-c935-b308-9b1ee2bf78c8/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/27efbfff-d66c-c935-b308-9b1ee2bf78c8/preview.webp",
					CreatedAt:  time.Date(2021, 2, 26, 19, 0, 0, 0, loc),
					Duration:   time.Hour + 12*time.Minute + 33*time.Second,
					Public:     true,
					Listed:     true,
					Tags:       []*Tag{{ID: TestTagDiskussionID}},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataRecording5Lang1,
					RecordingID: testdataRecording5,
					Lang:        "de",
					Title:       "Geschichte wird gemacht",
					Subtitle:    "aber wie und von wem?",
					Short:       "Das Erinnern an die Verbrechen und ein würdiges Gedenken an die Opfer des Nationalsozialismus ist und bleibt wichtig. Darin sind sich ein großer Teil der Bremer:innen und auch viele Politiker:innen einig. Doch wie sollten diese Erinnerung(en) eigentlich gestaltet sein? ...",
					Long: `
Das Erinnern an die Verbrechen und ein würdiges Gedenken an die Opfer des Nationalsozialismus ist und bleibt wichtig. Darin sind sich ein großer Teil der Bremer:innen und auch viele Politiker:innen einig. Doch wie sollten diese Erinnerung(en) eigentlich gestaltet sein? Wie sollten die heutigen Kenntnisse und die geschichtlichen Perspektiven vermittelt werden? Wie können die Auseinandersetzungen um die historischen Orte in unserer unmittelbaren Umgebung aktuell bleiben? Diese und andere Fragen wollen wir als ein Zusammenschluß freier Mitarbeiter:innen am Denkort Bunker Valentin in einem digitalen Forum mit den Gästen diskutieren. Bewusst wollen wir damit die Konvention einer starren Frontalveranstaltung aufbrechen und miteinander ins Gespräch kommen. In einer anschließenden Podiumsdiskussion wollen wir dann abgleichen, wie es um die „gesellschaftliche Verantwortung“ in Bremen und anderswo praktisch bestellt ist. Wer macht diese Arbeit und unter welchen Bedingungen? Wir diskutieren mit einer Aktiven des „Arisierungs“-Mahnmals in Bremen und der Initiative „Geschichte wird gemacht“ aus Berlin. Gemeinsam sollen Grenzen und Chancen einer Geschichtsvermittlung diskutiert werden, die eine lebendige Erinnerungskultur nicht nur beredet sondern umsetzt. Alle Interessierten sind herzlich willkommen!

Eine Veranstaltung von Erinnerungskultur anstellen. Organisierung freier Mitarbeitender am Denkort Bunker Valentin in kooperation mit dem Kulturzentrum Kukoon und Erinnern für die Zukunft e.V. sowie der Rosa-Luxemburg-Initiative – Die Rosa-Luxemburg-Stiftung in Bremen. 
				`,
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingFormat{
					ID:          testdataRecording5Format1,
					RecordingID: testdataRecording5,
					Lang:        "de",
					Quality:     0,
					IsVideo:     true,
					URL:         "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/27efbfff-d66c-c935-b308-9b1ee2bf78c8/video_best.mp4",
					Bytes:       862470450,
					Resolution:  "1920x1080",
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&RecordingFormat{
					ID: testdataRecording5Format1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&RecordingLang{
					ID: testdataRecording5Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Recording{
					ID: testdataRecording5,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0020-01-recording-6",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataRecording6,
					ChannelID:  TestChannelID1,
					CommonName: "2021-03-verschwoerungserzaehlung",
					Poster:     "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/81b262e9-e010-1fa2-84a5-d8cee1a94835/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/81b262e9-e010-1fa2-84a5-d8cee1a94835/preview.webp",
					CreatedAt:  time.Date(2021, 3, 3, 19, 0, 0, 0, loc),
					Duration:   time.Hour + 14*time.Minute + 17*time.Second,
					Public:     true,
					Listed:     true,
					Tags:       []*Tag{{ID: TestTagVortragID}, {ID: TestTagDiskussionID}},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataRecording6Speaker1,
							Name:    "Johanna Bröse",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataRecording6Speaker2,
							Name:    "Andrea Strübe",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataRecording6Lang1,
					RecordingID: testdataRecording6,
					Lang:        "de",
					Title:       "Welche Funktion haben Verschwörungserzählungen?",
					Subtitle:    "Warum der Glaube an einen Kapitalismus mit menschlichem Antlitz letztlich die größte Verschwörungsideologie ist.",
					Short:       "Verschwörungen – es gibt sie wirklich. Sie sind ein wichtiges Instrument zur Sicherung der politischen und gesellschaftlichen Macht in der Klassengesellschaft, aber auch bei Machtkämpfen unterschiedlicher Interessensgruppen untereinander oder im Kampf gegen Systemalternativen. Der Kapitalismus als Klassengesellschaft ist darauf angewiesen, den Antagonismus zwischen Kapitalisten und Lohnabhängigen aufrecht zu erhalten, und die Arbeiter*innenklasse auch durch Strategien der Verschleierung der Ausbeutungsverhältnisse von einer weitreichenden Organisierung abzuhalten. ...",
					Long: `Verschwörungen – es gibt sie wirklich. Sie sind ein wichtiges Instrument zur Sicherung der politischen und gesellschaftlichen Macht in der Klassengesellschaft, aber auch bei Machtkämpfen unterschiedlicher Interessensgruppen untereinander oder im Kampf gegen Systemalternativen. Der Kapitalismus als Klassengesellschaft ist darauf angewiesen, den Antagonismus zwischen Kapitalisten und Lohnabhängigen aufrecht zu erhalten, und die Arbeiter\*innenklasse auch durch Strategien der Verschleierung der Ausbeutungsverhältnisse von einer weitreichenden Organisierung abzuhalten. Viele der realen Verschwörungen wurden früher oder später aufgedeckt – durch kritische Journalist\*innen, Forscher\*innen, Aktivist\*innen.  Wiederum andere Verschwörungserzählungen – wie die, dass US-Eliten einen grausamen Handel mit Kindern aus einer Pizzeria heraus organisierten, konnten nie bewiesen werden. Warum? Weil sie schlicht und ergreifend falsch sind. Es gibt natürlich von vielen Menschen ein berechtigtes Unbehagen bis hin zum offenen Widerstand gegenüber dem gesellschaftlichen System, in dem wir leben. Ausbeutung, Ungleichheit, Klassenverhältnisse, struktureller Rassismus und Sexismus – um nur einige zu nennen – sind Erscheinungen eines globalen Kapitalismus. Aber nicht nur fortschrittliche Linke haben diesem den Kampf angesagt. Die Wut der Anhänger\*innen von Verschwörungserzählungen richtet sich gegen „die Multimilliardäre“, „das Establishment“ oder gegen „die Regierung“ –  ihre Macht wird aber mit einem Potpourri aus antisemitischen, rassistischen, antikommunistischen, antifeministischen und öfter auch esoterisch-wissenschaftsfeindlichen Versatzstücken erklärt. Verschwörungsanhänger\*innen versuchen also, grob gesagt, reale politische und gesellschaftliche Konflikte durch Machenschaften einer geheimen Gruppe zu erklären. Wie aber sollte man diesen Theorien und ihren Anhänger\*innen begegnen? Wie hängen Verschwörungstheorien und rechte Gesinnung zusammen? Und wie können wir produktiv mit der Erkenntnis umgehen, dass ein Kapitalismus mit menschlichem Antlitz letztlich die virulenteste Verschwörungserzählung ist?

Eine Veranstaltung von [kritisch-lesen.de](https://kritisch-lesen.de) in Kooperation mit dem Kulturzentrum Kukoon.
				`,
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingFormat{
					ID:          testdataRecording6Format1,
					RecordingID: testdataRecording6,
					Lang:        "de",
					Quality:     0,
					IsVideo:     true,
					URL:         "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/81b262e9-e010-1fa2-84a5-d8cee1a94835/video_best.mp4",
					Bytes:       1426234816,
					Resolution:  "1920x1080",
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&RecordingFormat{
					ID: testdataRecording6Format1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&RecordingLang{
					ID: testdataRecording6Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Recording{
					ID: testdataRecording6,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0020-01-recording-7",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataRecording7,
					ChannelID:  TestChannelID1,
					CommonName: "2021-03-kriegsgefanngende_in_bremen",
					Poster:     "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/728edaf7-9ad9-f972-4d09-ba5940cd43f9/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/728edaf7-9ad9-f972-4d09-ba5940cd43f9/preview.webp",
					CreatedAt:  time.Date(2021, 3, 4, 19, 0, 0, 0, loc),
					Duration:   time.Hour + 4*time.Minute + 25*time.Second,
					Public:     true,
					Listed:     true,
					Tags:       []*Tag{{ID: TestTagVortragID}},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataRecording7Speaker1,
							Name:    "Andreas Ehresmann",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataRecording7Speaker2,
							Name:    "Ronald Sperling",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataRecording7Speaker3,
							Name:    "Ines Dirolf",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataRecording7Lang1,
					RecordingID: testdataRecording7,
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
				if err := tx.Create(&RecordingFormat{
					ID:          testdataRecording7Format1,
					RecordingID: testdataRecording7,
					Lang:        "de",
					Quality:     0,
					IsVideo:     true,
					URL:         "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/728edaf7-9ad9-f972-4d09-ba5940cd43f9/video_best.mp4",
					Bytes:       958856106,
					Resolution:  "1920x1080",
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&RecordingFormat{
					ID: testdataRecording7Format1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&RecordingLang{
					ID: testdataRecording7Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Recording{
					ID: testdataRecording7,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0020-01-recording-8",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataRecording8,
					ChannelID:  TestChannelID1,
					CommonName: "2021-03-kriegsgefanngende_in_bremen",
					Poster:     "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/06e3a71e-581d-4735-9647-3e4a49b5caa8/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/06e3a71e-581d-4735-9647-3e4a49b5caa8/preview.webp",
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
							ID:      testdataRecording8Speaker1,
							Name:    "Andreas Speit",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataRecording8Lang1,
					RecordingID: testdataRecording8,
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
				if err := tx.Create(&RecordingFormat{
					ID:          testdataRecording8Format1,
					RecordingID: testdataRecording8,
					Lang:        "de",
					Quality:     0,
					IsVideo:     true,
					URL:         "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/06e3a71e-581d-4735-9647-3e4a49b5caa8/video_best.mp4",
					Bytes:       995280142,
					Resolution:  "1920x1080",
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&RecordingFormat{
					ID: testdataRecording8Format1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&RecordingLang{
					ID: testdataRecording8Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Recording{
					ID: testdataRecording8,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0020-01-recording-9",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataRecording9,
					ChannelID:  TestChannelID1,
					CommonName: "2021-04_nsu-watch",
					CreatedAt:  time.Date(2021, 4, 22, 19, 0, 0, 0, loc),
					Poster:     "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/57de7dfd-c060-4da1-8f57-f0880c1f2e5e/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/57de7dfd-c060-4da1-8f57-f0880c1f2e5e/preview.webp",
					Duration:   2*time.Hour + 18*time.Minute + 41*time.Second,
					Public:     true,
					Listed:     true,
					Tags: []*Tag{
						{ID: TestTagBuchvorstellungID},
						{ID: TestTagDiskussionID},
					},
					Speakers: []*Speaker{
						{
							OwnerID:      TestChannelID1,
							ID:           testdataRecording9Speaker1,
							Name:         "Caro Keller",
							Organisation: "NSU-Watch",
						},
						{
							OwnerID:      TestChannelID1,
							ID:           testdataRecording9Speaker2,
							Name:         "Lee Hielscher",
							Organisation: "Initiative in Gedenken an Nguyễn Ngọc Châu und Đỗ Anh Lân",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataRecording9Lang1,
					RecordingID: testdataRecording9,
					Lang:        "de",
					Title:       "Aufklären und Einmischen",
					Subtitle:    "Der NSU-Komplex und der Münchner Prozess - Buchvorstellung mit NSU-Watch",
					Short:       `Im November 2011 kam eine rechtsterroristische Mord- und Anschlagsserie des sogenannten Nationalsozialistischen Untergrunds (NSU) ans Licht, die in ihrer Dimension neu war. In den folgenden Untersuchungen formte sich ein erstes Bild des NSU-Komplexes. ...`,
					Long: `Im November 2011 kam eine rechtsterroristische Mord- und Anschlagsserie des sogenannten Nationalsozialistischen Untergrunds (NSU) ans Licht, die in ihrer Dimension neu war. In den folgenden Untersuchungen formte sich ein erstes Bild des NSU-Komplexes. Dabei wurde deutlich, dass eine noch umfassendere juristische und gesellschaftliche Aufarbeitung anstand. So beschlossen antifaschistische Initiativen und Einzelpersonen, die Arbeit am NSU-Komplex zu verstetigen, und gründeten »NSU-Watch«. Neun Jahre später ist die Aufarbeitung des NSU-Komplexes noch lange nicht abgeschlossen, die Gefahr des rechten Terrors bleibt schrecklich aktuell. NSU-Watch hat den NSU-Prozess beobachtet, jeden Tag protokolliert und der Öffentlichkeit zur Verfügung gestellt. Darüber hinaus haben sich Landesprojekte gegründet, die die parlamentarischen Aufklärungsbemühungen begleiten. Das zentrale Anliegen des Buches von NSU-Watch ist, die rassistischen Strukturen, die den NSU hervorbrachten, ihn wissentlich oder unwissentlich unterstützten und so zehn Morde, drei Sprengstoffanschläge und 15 Raubüberfälle zwischen 1998 und 2011 möglich machten, entlang der Geschehnisse und Akteur*innen des NSU-Prozesses in München aufzuzeigen. Trotz der vielen offen gebliebenen Fragen soll das Buch eine Zwischenbilanz bieten, die antifaschistischer Demokratieförderung zugrunde gelegt werden kann.

**NSU-Watch**

Das Autor\*innen-Kollektiv NSU-WATCH besteht aus Mitgliedern der unabhängigen Beobachtungsstelle NSU-Watch – Aufklären & Einmischen, die sich im Jahr 2012 gegründet hat, um die Aufklärungsbemühungen zum NSU-Komplex zu unterstützen und kritisch zu begleiten. NSU-Watch wird von einem Bündnis aus rund einem Dutzend antifaschistischer und antirassistischer Gruppen und Einzelpersonen aus dem ganzen Bundesgebiet getragen, die teilweise seit Jahrzehnten zum Themenkomplex Rechter Terror arbeiten. Kern der Arbeit von NSU-Watch war bzw. ist die Beobachtung des NSU-Prozesses am Oberlandesgericht in München sowie der diversen parlamentarischen Untersuchungsausschüsse im Bundestag und in den Ländern.

Eine Veranstaltung des _Kulturzentrum Kukoon_ in Kooperation mit der _Rosa-Luxemburg-Initiative – Die Rosa-Luxemburg-Stiftung in Bremen_.
				`,
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingFormat{
					ID:          testdataRecording9Format1,
					RecordingID: testdataRecording9,
					Lang:        "de",
					Quality:     0,
					IsVideo:     true,
					URL:         "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/57de7dfd-c060-4da1-8f57-f0880c1f2e5e/video_best.mp4",
					Bytes:       2186863051,
					Resolution:  "1920x1080",
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&RecordingFormat{
					ID: testdataRecording9Format1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&RecordingLang{
					ID: testdataRecording9Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Recording{
					ID: testdataRecording9,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0020-01-recording-10",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataRecording10,
					ChannelID:  TestChannelID1,
					EventID:    &TestEventID1,
					CommonName: "2021-05_out_loud-mareice_kaiser-modernen_mutter",
					CreatedAt:  time.Date(2021, 5, 5, 19, 0, 0, 0, loc),
					Poster:     "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/4fb029d6-063a-4302-9ae8-4c1c6a1542a5/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/4fb029d6-063a-4302-9ae8-4c1c6a1542a5/preview.webp",
					Duration:   time.Hour + 28*time.Minute + 26*time.Second,
					Public:     true,
					Listed:     true,
					Tags: []*Tag{
						{ID: TestTagBuchvorstellungID},
						{ID: TestTagDiskussionID},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataRecording10Lang1,
					RecordingID: testdataRecording10,
					Lang:        "de",
					Title:       "Mareice Kaiser",
					Subtitle:    "Das Unwohlsein der modernen Mutter",
					Short:       `Mütter sollen heute alles sein: Versorgerin, Businesswoman, Mom I'd like to fuck. Dass darunter ihr Wohlbefinden leidet, ist kein Wunder. Mareice Kaiser, Journalistin und selbst Mutter, stellt dabei immer wieder fest: ...`,
					Long:        `Mütter sollen heute alles sein: Versorgerin, Businesswoman, Mom I'd like to fuck. Dass darunter ihr Wohlbefinden leidet, ist kein Wunder. Mareice Kaiser, Journalistin und selbst Mutter, stellt dabei immer wieder fest: Das Mutterideal ist unerreichbar und voller Widersprüche. Nichts kann man richtig machen und niemandem etwas recht. Mutterschaft berührt dabei, natürlich, jeden Lebensbereich: Denn egal, ob es um Arbeit, Geld, Sex, Körper, Psyche oder Liebe geht – Stereotype, Klischees und gesellschaftlichen Druck gibt es überall, auf Instagram, im Bett und im Büro. In ihrem Buch "Das Unwohlsein der modernen Mutter" (Rowohlt, 2021) zeigt die Autorin, wo Mütter heute stehen: noch immer öfter am Herd als in den Chefetagen. Und, wo sie stehen sollten: Dort, wo sie selbst sich sehen – frei und selbstbestimmt. Bei OUT LOUD liest Mareice Kaiser aus ihrem Buch und spricht mit uns über Frausein, Mutterschaft und Selbstbestimmung.`,
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingFormat{
					ID:          testdataRecording10Format1,
					RecordingID: testdataRecording10,
					Lang:        "de",
					Quality:     0,
					IsVideo:     true,
					URL:         "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/4fb029d6-063a-4302-9ae8-4c1c6a1542a5/video_best.mp4",
					Bytes:       2443666130,
					Resolution:  "1920x1080",
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&RecordingFormat{
					ID: testdataRecording10Format1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&RecordingLang{
					ID: testdataRecording10Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Recording{
					ID: testdataRecording10,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0030-01-stream-1",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream1,
					ChannelID: TestChannelID1,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 3, 4, 18, 30, 0, 0, loc),
					ListenAt:  time.Date(2021, 3, 4, 18, 30, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/728edaf7-9ad9-f972-4d09-ba5940cd43f9/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/728edaf7-9ad9-f972-4d09-ba5940cd43f9/preview.webp",
					Tags:      []*Tag{{ID: TestTagVortragID}},
					Speakers: []*Speaker{
						{
							ID: testdataRecording7Speaker1,
						},
						{
							ID: testdataRecording7Speaker2,
						},
						{
							ID: testdataRecording7Speaker3,
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStream1Lang1,
					StreamID: testdataStream1,
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
				if err := tx.Delete(&StreamLang{
					ID: testdataStream1Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Stream{
					ID: testdataStream1,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0030-01-stream-2",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream2,
					ChannelID: TestChannelID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 3, 4, 18, 30, 0, 0, loc),
					ListenAt:  time.Date(2021, 3, 4, 18, 30, 0, 0, loc),
					Poster:    "https://c3woc.de/images/banner.jpg",
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&Stream{
					ID: testdataStream2,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0030-01-stream-3",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream3,
					ChannelID: TestChannelID1,
					Chat:      true,
					Running:   true,
					StartAt:   time.Date(2021, 4, 15, 0, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 4, 15, 0, 0, 0, 0, loc),
					Poster:    "https://media.kukoon.de/images/41185fa4-3e22-44bb-9020-1d824e12ede3.jpg",
					Tags: []*Tag{
						{ID: TestTagBuchvorstellungID},
						{ID: TestTagDiskussionID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataStream3Speaker1,
							Name:    "Andreas Speit",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStream3Lang1,
					StreamID: testdataStream3,
					Lang:     "de",
					Title:    "Rechte Egoshooter",
					Subtitle: "Von der virtuellen Hetze zum Livestream-Attentat",
					Short:    `Weltweit gibt es rechtsterroristische Attentate eines neuen Typs. In Halle (Saale) verhinderte nur eine verschlossene Holztür der Synagoge ein größeres Massaker. ...`,
					Long: `Weltweit gibt es rechtsterroristische Attentate eines neuen Typs. In Halle (Saale) verhinderte nur eine verschlossene Holztür der Synagoge ein größeres Massaker. Am 9. Oktober 2019 wollte dort ein Rechtsextremist die versammelten Juden hinrichten. Mit selbstgebauten Waffen schoss er auf die Tür und warf eigens hergestellte Sprengsätze. Online konnten Gleichgesinnte zusehen, wie er zwei Menschen ermordete: Seine Tat verbreitete er per Videokamera auf einem Portal für Computerspiel-Videos. Er ahmte damit andere »Egoshooter« nach – wie einen Rechtsextremisten, der in Neuseeland wenige Monate zuvor die Tötung von 51 Menschen live im Internet übertragen hatte. Was treibt Menschen vom Bildschirm zur realen Gewalt auf der Straße? Die Beiträge des Buches gehen den Spuren der Attentäter nach und zeigen die speziellen Radikalisierungsmechanismen im Netz auf. Sie erklären die Hintergründe und Motive dieser Männer, die in ihren rechten Online-Gemeinden Antisemitismus, Rassismus und Antifeminismus verbreiten. Das Buch gibt Einblicke in eine Welt, die vielen unbekannt ist.

Eine Veranstaltung des _Kulturzentrum Kukoon_ in Kooperation mit der _Rosa-Luxemburg-Initiative – Die Rosa-Luxemburg-Stiftung in Bremen_.
				`,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&StreamLang{
					ID: testdataStream3Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Stream{
					ID: testdataStream3,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0030-01-stream-4",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream4,
					ChannelID: TestChannelID1,
					Chat:      true,
					Running:   true,
					StartAt:   time.Date(2021, 4, 22, 0, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 4, 22, 0, 0, 0, 0, loc),
					Poster:    "https://media.kukoon.de/images/f61fa2de-72d3-4a1e-98b2-65b13d8ecb01.jpg",
					Tags: []*Tag{
						{ID: TestTagBuchvorstellungID},
						{ID: TestTagDiskussionID},
					},
					Speakers: []*Speaker{
						{
							ID: testdataStream4Speaker1,
						},
						{
							ID: testdataStream4Speaker2,
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStream4Lang1,
					StreamID: testdataStream4,
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
				if err := tx.Delete(&StreamLang{
					ID: testdataStream4Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Stream{
					ID: testdataStream4,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0030-01-stream-5",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream5,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID1,
					Chat:      true,
					Running:   true,
					StartAt:   time.Date(2021, 5, 5, 0, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 5, 5, 0, 0, 0, 0, loc),
					Poster:    "https://media.kukoon.de/images/67bd5c4c-81d6-47c8-adb9-458a9da58dbd.jpg",
					Tags: []*Tag{
						{ID: TestTagBuchvorstellungID},
						{ID: TestTagDiskussionID},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStream5Lang1,
					StreamID: testdataStream5,
					Lang:     "de",
					Title:    "Mareice Kaiser",
					Subtitle: "Das Unwohlsein der modernen Mutter",
					Short:    `Mütter sollen heute alles sein: Versorgerin, Businesswoman, Mom I'd like to fuck. Dass darunter ihr Wohlbefinden leidet, ist kein Wunder. Mareice Kaiser, Journalistin und selbst Mutter, stellt dabei immer wieder fest: ...`,
					Long:     `Mütter sollen heute alles sein: Versorgerin, Businesswoman, Mom I'd like to fuck. Dass darunter ihr Wohlbefinden leidet, ist kein Wunder. Mareice Kaiser, Journalistin und selbst Mutter, stellt dabei immer wieder fest: Das Mutterideal ist unerreichbar und voller Widersprüche. Nichts kann man richtig machen und niemandem etwas recht. Mutterschaft berührt dabei, natürlich, jeden Lebensbereich: Denn egal, ob es um Arbeit, Geld, Sex, Körper, Psyche oder Liebe geht – Stereotype, Klischees und gesellschaftlichen Druck gibt es überall, auf Instagram, im Bett und im Büro. In ihrem Buch "Das Unwohlsein der modernen Mutter" (Rowohlt, 2021) zeigt die Autorin, wo Mütter heute stehen: noch immer öfter am Herd als in den Chefetagen. Und, wo sie stehen sollten: Dort, wo sie selbst sich sehen – frei und selbstbestimmt. Bei OUT LOUD liest Mareice Kaiser aus ihrem Buch und spricht mit uns über Frausein, Mutterschaft und Selbstbestimmung.`,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&StreamLang{
					ID: testdataStream5Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Stream{
					ID: testdataStream5,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0030-01-stream-6",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream6,
					ChannelID: TestChannelID1,
					Chat:      true,
					Running:   true,
					StartAt:   time.Date(2021, 6, 24, 0, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 6, 24, 0, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/1742d9b6-c9c6-45fb-a3a3-4a3e7fac2987/poster.png",
					Tags: []*Tag{
						{ID: TestTagDiskussionID},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStream6Lang1,
					StreamID: testdataStream6,
					Lang:     "de",
					Title:    "System Change not Climate Change!",
					Subtitle: "Einführung zu Klimakrise und Kapitalismuskritik",
					Short: `Diskussionsveranstaltung mit der Gruppe **direction f** (Hannover)

**Kukoon im Park** oder hier`,
					Long: `Diskussionsveranstaltung mit der Gruppe **direction f** (Hannover)

**Kukoon im Park** oder hier

Zwar verblasst die Klimakrise seit einem Jahr im medialen Schatten der Corona-Pandemie, die Dringlichkeit zum Handeln bleibt jedoch unverändert. Klar ist, dass die Klimakrise kein rein ökologisches Phänomen sondern ebenso sehr eine soziale Krise ist. Als »direction f« haben wir uns bisher vorrangig mit den Zusammenhängen von Klimakrise und Kapitalismus befasst. Im Rahmen der Veranstaltung wollen wir kurz auf den Ist-Zustand und bestehende Zusammenhänge eingehen. Davon ausgehend würden wir gerne darüber diskutieren, was (un)taugliche Strategien gegen die drohende Klimakatstrophe sein können und welche Rolle und Aufgaben dabei einer (radikalen) Linken zukämen. direction f ist ein Zusammenschluss von Menschen in Hannover, der sich bisher schwerpunktmäßig mit dem Zusammenhang von Klimakrise und Kapitalismus befasst hat.

Mehr Infos unter [direction-f.org](https://direction-f.org/)`,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&StreamLang{
					ID: testdataStream6Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Stream{
					ID: testdataStream6,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
	}...)
}
