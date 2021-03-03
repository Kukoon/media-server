package models

import (
	"time"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	loc = time.FixedZone("UTC+2", +2*60*60)

	testdataChannel1 = uuid.MustParse("df1555f5-7046-4f7a-adcc-195b73949723")

	testdataStream1         = uuid.MustParse("dffe2c0e-3713-4399-8ee2-279becbbb06e")
	testdataStream1Lang1    = uuid.MustParse("3a4f9157-65bf-4d15-a82b-1cd9295d07e0")
	testdataStream1Speaker1 = uuid.MustParse("0d1b38cd-561c-4db4-b4b9-51f74ba3dba4")
	testdataStream1Speaker2 = uuid.MustParse("1dbf0438-a9c1-4412-b44c-08fe7819902c")
	testdataStream1Speaker3 = uuid.MustParse("d68e5de7-e56e-46a7-843c-4a06e540cf3a")

	testdataTagBuchvorstellung     = uuid.MustParse("0bca0cf4-a9b9-46d7-821f-18c59c08fc1d")
	testdataTagBuchvorstellungLang = uuid.MustParse("35822fe2-1910-48e7-904f-15c9e6f7ea34")
	testdataTagDiskussion          = uuid.MustParse("277026b0-b9d6-48d6-bfa1-96dcc7eb3451")
	testdataTagDiskussionLang      = uuid.MustParse("38722845-beba-4e3d-ad3f-694c029d751f")
	testdataTagVortrag             = uuid.MustParse("7297a654-71f9-43be-8120-69b8152f01fc")
	testdataTagVortragLang         = uuid.MustParse("ec784c8e-2673-4870-b219-eb636e4765c8")

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
)

var testdata = []*gormigrate.Migration{
	{
		ID: "10-data-0010-01-channel",
		Migrate: func(tx *gorm.DB) error {
			return tx.Create(&Channel{
				ID:         testdataChannel1,
				CommonName: "kukoon",
				Title:      "Im Kukoon",
				Logo:       "https://media.kukoon.de/static/css/kukoon/logo.png",
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Delete(&Recording{
				ID: uuid.MustParse("df1555f5-7046-4f7a-adcc-195b73949723"),
			}).Error
		},
	},
	{
		ID: "10-data-0018-01-tag-eg",
		Migrate: func(tx *gorm.DB) error {
			// -
			if err := tx.Create(&Tag{
				ID: testdataTagBuchvorstellung,
			}).Error; err != nil {
				return err
			}
			if err := tx.Create(&TagLang{
				ID:    testdataTagBuchvorstellungLang,
				TagID: testdataTagBuchvorstellung,
				Lang:  "de",
				Name:  "Buchvorstellung",
			}).Error; err != nil {
				return err
			}
			// -
			if err := tx.Create(&Tag{
				ID: testdataTagDiskussion,
			}).Error; err != nil {
				return err
			}
			if err := tx.Create(&TagLang{
				ID:    testdataTagDiskussionLang,
				TagID: testdataTagDiskussion,
				Lang:  "de",
				Name:  "Diskussion",
			}).Error; err != nil {
				return err
			}
			// -
			if err := tx.Create(&Tag{
				ID: testdataTagVortrag,
			}).Error; err != nil {
				return err
			}
			if err := tx.Create(&TagLang{
				ID:    testdataTagVortragLang,
				TagID: testdataTagVortrag,
				Lang:  "de",
				Name:  "Vortrag",
			}).Error; err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// -
			if err := tx.Delete(&TagLang{ID: testdataTagVortragLang}).Error; err != nil {
				return err
			}
			if err := tx.Delete(&Tag{ID: testdataTagVortrag}).Error; err != nil {
				return err
			}
			// -
			if err := tx.Delete(&TagLang{ID: testdataTagDiskussionLang}).Error; err != nil {
				return err
			}
			if err := tx.Delete(&Tag{ID: testdataTagDiskussion}).Error; err != nil {
				return err
			}
			// -
			if err := tx.Delete(&TagLang{ID: testdataTagBuchvorstellungLang}).Error; err != nil {
				return err
			}
			if err := tx.Delete(&Tag{ID: testdataTagBuchvorstellung}).Error; err != nil {
				return err
			}
			return nil
		},
	},
	{
		ID: "10-data-0020-01-recording-1",
		Migrate: func(tx *gorm.DB) error {
			if err := tx.Create(&Recording{
				ID:         testdataRecording1,
				ChannelID:  testdataChannel1,
				CommonName: "2020-12-polizeigewalt",
				Poster:     "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee/poster.png",
				Preview:    "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee/preview.gif",
				CreatedAt:  time.Date(2020, 12, 10, 20, 0, 0, 0, loc),
				Duration:   time.Hour,
				Public:     true,
				Listed:     true,
				Tags:       []*Tag{{ID: testdataTagDiskussion}},
				Speakers: []*Speaker{
					{
						OwnerID:      testdataChannel1,
						ID:           testdataRecording1Speaker1,
						Name:         "Leila Abdul-Rahman",
						Organisation: "Ruhr Universität Bochum",
					},
					{
						OwnerID:      testdataChannel1,
						ID:           testdataRecording1Speaker2,
						Name:         "Greta",
						Organisation: "Grün-Weiße Hilfe",
					},
					{
						OwnerID:      testdataChannel1,
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
				Title:       "Polizeigewalt – ein deutsches Problem",
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
				URL:         "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee/video_best.mp4",
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
				URL:         "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee/video_720.mp4",
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
				URL:         "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee/video_480.mp4",
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
				URL:         "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee/audio_best.mp3",
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
				ChannelID:  testdataChannel1,
				CommonName: "2021-01-faschistische_jahrhundert",
				Poster:     "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00/poster.png",
				Preview:    "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00/preview.gif",
				CreatedAt:  time.Date(2021, 01, 29, 20, 0, 0, 0, loc),
				Duration:   time.Hour,
				Public:     true,
				Listed:     true,
				Tags: []*Tag{
					{ID: testdataTagBuchvorstellung},
					{ID: testdataTagDiskussion},
				},
				Speakers: []*Speaker{
					{
						OwnerID: testdataChannel1,
						ID:      testdataRecording2Speaker1,
						Name:    "Volkmar Wölk",
					},
					{
						OwnerID: testdataChannel1,
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
				URL:         "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00/video_best.mp4",
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
				URL:         "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00/video_720.mp4",
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
				URL:         "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00/video_480.mp4",
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
				URL:         "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00/audio_best.mp3",
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
				ChannelID:  testdataChannel1,
				CommonName: "2021-02-pushbacks",
				Poster:     "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/edb1cfbb-3476-d639-b3f5-795fabf4ef4d/poster.png",
				Preview:    "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/edb1cfbb-3476-d639-b3f5-795fabf4ef4d/preview.gif",
				CreatedAt:  time.Date(2021, 2, 5, 20, 0, 0, 0, loc),
				Duration:   time.Hour,
				Public:     true,
				Listed:     true,
				Tags:       []*Tag{{ID: testdataTagVortrag}},
				Speakers: []*Speaker{
					{
						OwnerID: testdataChannel1,
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
				URL:         "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/edb1cfbb-3476-d639-b3f5-795fabf4ef4d/video_best.mp4",
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
		ID: "10-data-0030-01-stream-1",
		Migrate: func(tx *gorm.DB) error {
			if err := tx.Create(&Stream{
				ID:        testdataStream1,
				ChannelID: testdataChannel1,
				Chat:      false,
				Running:   true,
				Poster:    "https://media.kukoon.de/images/00ff627e-cb44-4186-acf3-e7d38d63db74.jpg",
				Tags:      []*Tag{{ID: testdataTagVortrag}},
				Speakers: []*Speaker{
					{
						OwnerID: testdataChannel1,
						ID:      testdataStream1Speaker1,
						Name:    "Andreas Ehresmann",
					},
					{
						OwnerID: testdataChannel1,
						ID:      testdataStream1Speaker2,
						Name:    "Ronald Sperling",
					},
					{
						OwnerID: testdataChannel1,
						ID:      testdataStream1Speaker3,
						Name:    "Ines Dirolf",
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
}
