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

	testdataRecording1        = uuid.MustParse("542685cb-3693-e720-a957-f008f5dae3ee")
	testdataRecording1Lang1   = uuid.MustParse("03d33e6a-151f-47d9-be79-a726e0f9a859")
	testdataRecording1Format1 = uuid.MustParse("3a4f9157-65bf-4d15-a82b-1cd9295d07e0")

	testdataRecording2        = uuid.MustParse("45da89a7-e5e0-5104-b937-6d4c2d4b6d00")
	testdataRecording2Lang1   = uuid.MustParse("20eb49e4-46a6-4b30-87eb-72e4dbe77b09")
	testdataRecording2Format1 = uuid.MustParse("09858a43-0532-4ad8-8694-92ed71372ff4")
	testdataRecording2Format2 = uuid.MustParse("6b4cf4b0-db39-47a3-9bc9-4ff6ef539a88")
	testdataRecording2Format3 = uuid.MustParse("918c0386-03ad-4d34-b930-d6ce6a7632eb")

	testdataRecording3        = uuid.MustParse("edb1cfbb-3476-d639-b3f5-795fabf4ef4d")
	testdataRecording3Lang1   = uuid.MustParse("d193bc41-99f6-46d8-870e-72a860520223")
	testdataRecording3Format1 = uuid.MustParse("6b1b95f2-d92d-4da7-b56c-1ba86ff22dcd")
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
				Secret:     uuid.MustParse("1f349cf3-196d-4e39-9d22-3e35497e990c"),
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Delete(&Recording{
				ID: uuid.MustParse("df1555f5-7046-4f7a-adcc-195b73949723"),
			}).Error
		},
	},
	{
		ID: "10-data-0020-01-recording-1",
		Migrate: func(tx *gorm.DB) error {
			err := tx.Create(&Recording{
				ID:         testdataRecording1,
				ChannelID:  testdataChannel1,
				CommonName: "2020-12-polizeigewalt",
				Poster:     "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee/poster.png",
				Preview:    "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee/preview.gif",
				CreatedAt:  time.Date(2020, 12, 10, 20, 0, 0, 0, loc),
				Duration:   time.Hour,
				Public:     true,
				Listed:     true,
			}).Error
			if err != nil {
				return err
			}
			err = tx.Create(&RecordingLang{
				ID:          testdataRecording1Lang1,
				RecordingID: testdataRecording1,
				Lang:        "de",
				Title:       "Polizeigewalt – ein deutsches Problem",
				Description: `Ein deutsches Problem Diskussionsveranstaltung mit Laila Abdul-Rahman, Greta ([Grün-Weiße Hilfe Bremen](https://twitter.com/fanhilfe_bremen?lang=de)) und Mathilda ([Kampagne für Opfer rassistischer Polizeigewalt – KOP Bremen](https://www.facebook.com/KOP-Bremen-Kampagne-f%C3%BCr-Opfer-rassistischer-Polizeigewalt-Bremen-168776953679814/))

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
			}).Error
			if err != nil {
				return err
			}
			return tx.Create(&RecordingFormat{
				ID:          testdataRecording1Format1,
				RecordingID: testdataRecording1,
				Lang:        "de",
				Quality:     0,
				IsVideo:     true,
				URL:         "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee/video_best.mp4",
				Bytes:       3323919713,
				Resolution:  "1920x1080",
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.Delete(&RecordingFormat{
				ID: testdataRecording1Format1,
			}).Error
			if err != nil {
				return err
			}
			err = tx.Delete(&RecordingLang{
				ID: testdataRecording1Lang1,
			}).Error
			if err != nil {
				return err
			}
			return tx.Delete(&Recording{
				ID: testdataRecording1,
			}).Error
		},
	},
	{
		ID: "10-data-0020-01-recording-2",
		Migrate: func(tx *gorm.DB) error {
			err := tx.Create(&Recording{
				ID:         testdataRecording2,
				ChannelID:  testdataChannel1,
				CommonName: "2021-01-faschistische_jahrhundert",
				Poster:     "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00/poster.png",
				Preview:    "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00/preview.gif",
				CreatedAt:  time.Date(2021, 01, 29, 20, 0, 0, 0, loc),
				Duration:   time.Hour,
				Public:     true,
				Listed:     true,
			}).Error
			if err != nil {
				return err
			}
			err = tx.Create(&RecordingLang{
				ID:          testdataRecording2Lang1,
				RecordingID: testdataRecording2,
				Lang:        "de",
				Title:       "Das faschistische Jahrhundert",
				Description: `## Buchvorstellung: Das faschistische Jahrhundert

„Wer das Wesen des Faschismus erkennen will, muss zurück zu dessen Wurzeln“ – so der Faschismusforscher Zeev Sternhell. Gleiches gilt für die europäische Neue Rechte. Aus welchen Gründen ist sie wann und wo mit welchen Zielen entstanden? Wo liegen ihre geistigen Wurzeln? Betreiben die „neuen“ Rechten tatsächlich eine ideologische Erneuerung der extremen Rechten, oder handelt es sich lediglich um alten Wein in neuen Schläuchen? **Volkmar Wölk** und **Felix Schilk**, Mitautoren des Bandes _Das faschistische Jahrhundert_, veranschaulichen die Denkwege der Neuen Rechten anhand von zwei zentralen Themenfelder: Einerseits ihre Europakonzeptionen und andererseits ihre wirtschafts- und sozialpolitischen Vorstellungen.

Ein kurzer Beitrag zu dem Buch [hier](https://www.deutschlandfunkkultur.de/friedrich-burschel-das-faschistische-jahrhundert-schon.1270.de.html?dram:article_id=485508).

Eine Veranstaltung des _Kulturzentrum Kukoon_ in Kooperation mit der _Rosa-Luxemburg-Initiative – Die Rosa-Luxemburg-Stiftung in Bremen_.
				`,
			}).Error
			if err != nil {
				return err
			}
			err = tx.Create(&RecordingFormat{
				ID:          testdataRecording2Format1,
				RecordingID: testdataRecording2,
				Lang:        "de",
				Quality:     0,
				IsVideo:     true,
				URL:         "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00/video_best.mp4",
				Bytes:       2878429977,
				Resolution:  "1920x1080",
			}).Error
			if err != nil {
				return err
			}
			err = tx.Create(&RecordingFormat{
				ID:          testdataRecording2Format2,
				RecordingID: testdataRecording2,
				Lang:        "de",
				Quality:     160,
				IsVideo:     true,
				URL:         "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00/video_720p.mp4",
				Bytes:       115967040,
				Resolution:  "1280x720",
			}).Error
			if err != nil {
				return err
			}
			return tx.Create(&RecordingFormat{
				ID:          testdataRecording2Format3,
				RecordingID: testdataRecording2,
				Lang:        "de",
				Quality:     0,
				IsVideo:     false,
				URL:         "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00/audio_best.mp3",
				Bytes:       115967040,
				Resolution:  "149kb",
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.Delete(&RecordingFormat{
				ID: testdataRecording2Format3,
			}).Error
			if err != nil {
				return err
			}
			err = tx.Delete(&RecordingFormat{
				ID: testdataRecording2Format2,
			}).Error
			if err != nil {
				return err
			}
			err = tx.Delete(&RecordingFormat{
				ID: testdataRecording2Format1,
			}).Error
			if err != nil {
				return err
			}
			err = tx.Delete(&RecordingLang{
				ID: testdataRecording2Lang1,
			}).Error
			if err != nil {
				return err
			}
			return tx.Delete(&Recording{
				ID: testdataRecording2,
			}).Error
		},
	},
	{
		ID: "10-data-0020-01-recording-3",
		Migrate: func(tx *gorm.DB) error {
			err := tx.Create(&Recording{
				ID:         testdataRecording3,
				ChannelID:  testdataChannel1,
				CommonName: "2021-02-pushbacks",
				Poster:     "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/edb1cfbb-3476-d639-b3f5-795fabf4ef4d/poster.png",
				Preview:    "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/edb1cfbb-3476-d639-b3f5-795fabf4ef4d/preview.gif",
				CreatedAt:  time.Date(2021, 2, 5, 20, 0, 0, 0, loc),
				Duration:   time.Hour,
				Public:     true,
				Listed:     true,
			}).Error
			if err != nil {
				return err
			}
			err = tx.Create(&RecordingLang{
				ID:          testdataRecording3Lang1,
				RecordingID: testdataRecording3,
				Lang:        "de",
				Title:       "Pushbacks, Internierung, Zugangshürden – Zum Stand des europäischen Migrations- und Grenzregimes",
				Description: `Nach dem katastrophalen Brand des Flüchtlingslagers Moria auf Lesbos hatte die Europäische Kommission erneut einen Neustart in der europäischen Migrations- und Asylpolitik versucht. Der Neue Pakt für Asyl und Migration spitzt jedoch vor allem die Entwicklungen zu, die das europäische Asylsystem seit dem Sommer der Migration 2015 kennzeichnen: Zugangshürden zum Asylsystem und massive Internierung. Der Pakt adressiert jedoch nicht die bestehenden Rechtsverletzungen, die das europäische Grenzregime kennzeichnen: illegale Pushbacks sind an vielen Abschnitten der europäischen Außengrenze zur Normalität geworden, und scheinen von der europäischen Grenzschutzagentur Frontex stillschweigend geduldet, oder sogar aktiv unterstützt zu werden. Dies ist vor allem deswegen Besorgnis erregend, da die Agentur bis 2027 über ein eigenes Grenzschutzkorps mit 10.000 Grenzschützern verfügen soll. In der Veranstaltung werden all diese Entwicklungen des letzten Jahres skizziert werden, wobei eine Diskussion um Alternativen und Handlungsmöglichkeiten für ein solidarisches Europa von unten nicht zu kurz kommen sollen.

**Bernd Kasparek** ist seit vielen Jahren Aktivist und Grenzregimeforscher. Sein Buch „Europa als Grenze. Eine Ethnographie der Grenzschutz-Agentur Frontex“ wird im Mai bei transcript erscheinen. Er ist Mitglied des Netzwerks für Kritische Migrations- und Grenzregimeforschung und der Forschungsassoziation [bordermonitoring.eu](http://bordermonitoring.eu/) e.V.

Eine Veranstaltung des Kulturzentrum Kukoon in Kooperation mit der Rosa-Luxemburg-Initiative – Die Rosa-Luxemburg-Stiftung in Bremen.
				`,
			}).Error
			if err != nil {
				return err
			}
			return tx.Create(&RecordingFormat{
				ID:          testdataRecording3Format1,
				RecordingID: testdataRecording3,
				Lang:        "de",
				Quality:     0,
				IsVideo:     true,
				URL:         "https://v2.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/edb1cfbb-3476-d639-b3f5-795fabf4ef4d/video_best.mp4",
				Bytes:       1092701356,
				Resolution:  "1920x1080",
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.Delete(&RecordingFormat{
				ID: testdataRecording3Format1,
			}).Error
			if err != nil {
				return err
			}
			err = tx.Delete(&RecordingLang{
				ID: testdataRecording3Lang1,
			}).Error
			if err != nil {
				return err
			}
			return tx.Delete(&Recording{
				ID: testdataRecording3,
			}).Error
		},
	},
}
