package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	testdataID := uuid.MustParse("edb1cfbb-3476-d639-b3f5-795fabf4ef4d")
	testdataIDLang1 := uuid.MustParse("d193bc41-99f6-46d8-870e-72a860520223")
	testdataIDSpeaker1 := uuid.MustParse("a098c2f5-aa63-4c54-87b1-46ddda1cde16")
	testdataIDFormat1 := uuid.MustParse("6b1b95f2-d92d-4da7-b56c-1ba86ff22dcd")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-recording-03",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataID,
					ChannelID:  TestChannelID1,
					CommonName: "2021-02-pushbacks",
					Poster:     "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					CreatedAt:  time.Date(2021, 2, 5, 20, 0, 0, 0, loc),
					Duration:   time.Hour + 43*time.Minute + 21*time.Second,
					Public:     true,
					Listed:     true,
					Tags: []*Tag{
						{ID: TestTagVortragID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker1,
							Name:    "Bernd Kasparek",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
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
				if err := tx.Create(&[]*RecordingFormat{
					{
						ID:          testdataIDFormat1,
						RecordingID: testdataID,
						Lang:        "de",
						Quality:     0,
						IsVideo:     true,
						URL:         "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/video_best.mp4",
						Bytes:       1092701356,
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
