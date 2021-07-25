package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	testdataID := uuid.MustParse("45da89a7-e5e0-5104-b937-6d4c2d4b6d00")
	testdataIDLang1 := uuid.MustParse("20eb49e4-46a6-4b30-87eb-72e4dbe77b09")
	testdataIDSpeaker1 := uuid.MustParse("5dfddf98-2b11-48c7-ab5e-8cc6a2db0dcd")
	testdataIDSpeaker2 := uuid.MustParse("a85fffb1-aa9f-4bd0-81e7-0058891827ef")
	testdataIDFormat1 := uuid.MustParse("09858a43-0532-4ad8-8694-92ed71372ff4")
	testdataIDFormat2 := uuid.MustParse("6b4cf4b0-db39-47a3-9bc9-4ff6ef539a88")
	testdataIDFormat3 := uuid.MustParse("fc129d9b-b774-4b47-a7f3-a8cf046a5573")
	testdataIDFormat4 := uuid.MustParse("918c0386-03ad-4d34-b930-d6ce6a7632eb")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0020-01-recording-2",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataID,
					ChannelID:  TestChannelID1,
					CommonName: "2021-01-faschistische_jahrhundert",
					Poster:     "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
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
							ID:      testdataIDSpeaker1,
							Name:    "Volkmar Wölk",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker2,
							Name:    "Felix Schilk",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
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
				if err := tx.Create(&[]*RecordingFormat{
					{
						ID:          testdataIDFormat1,
						RecordingID: testdataID,
						Lang:        "de",
						Quality:     0,
						IsVideo:     true,
						URL:         "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/video_best.mp4",
						Bytes:       2878429977,
						Resolution:  "1920x1080",
					},
					{
						ID:          testdataIDFormat2,
						RecordingID: testdataID,
						Lang:        "de",
						Quality:     160,
						IsVideo:     true,
						URL:         "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/video_480.mp4",
						Bytes:       1149359246,
						Resolution:  "1280x720",
					},
					{
						ID:          testdataIDFormat3,
						RecordingID: testdataID,
						Lang:        "de",
						Quality:     180,
						IsVideo:     true,
						URL:         "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/video_480.mp4",
						Bytes:       480268045,
						Resolution:  "854x480",
					},
					{
						ID:          testdataIDFormat4,
						RecordingID: testdataID,
						Lang:        "de",
						Quality:     0,
						IsVideo:     false,
						URL:         "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/audio_best.mp4",
						Bytes:       115967040,
						Resolution:  "149kb",
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
