package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Grand Piano: Julia Kadel
	testdataID := uuid.MustParse("e446cfa4-9df2-4965-8491-b0d24a573585")
	testdataIDLang1 := uuid.MustParse("622f5adf-9d4d-4bef-8098-247c4180c6d7")
	testdataIDSpeaker1 := uuid.MustParse("e9d2cb0a-abc6-4ac8-ae8f-d462f87aad43")
	testdataIDFormat1 := uuid.MustParse("035fb6fb-1e89-4bc5-97f9-a0f63fffdf03")

	// see stream 17

	/* WARNING unreleased:
	- Public
	- Private
	- Duration
	- Bytes
	*/

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-recording-22",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					CreatedAt: time.Date(2021, 7, 25, 18, 0, 0, 0, loc),
					Duration:  time.Hour,
					Public:    false,
					Listed:    false,
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker1,
							Name:    "Julia Kadel",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
					Lang:        "de",
					Title:       "Julia Kadel",
					Subtitle:    "Sonntag ab 18:00",
					Short:       `**Julia Kadel** wurde 1986 in Berlin geboren und begann im Alter von sieben Jahren Klavier zu spielen. Nach Jahren der klassischen Ausbildung entdeckte sie mit Fünfzehn ihre Leidenschaft für den Jazz.`,
					Long: `**Julia Kadel** wurde 1986 in Berlin geboren und begann im Alter von sieben Jahren Klavier zu spielen. Nach Jahren der klassischen Ausbildung entdeckte sie mit Fünfzehn ihre Leidenschaft für den Jazz. Bis 2016 studierte sie Jazzklavier an der Hochschule für Musik Carl Maria von Weber Dresden. Seither zählt **Kadel** zu den starken neuen Stimmen innerhalb der europäischen Musikszene.
Im Frühjahr 2013 erhielt sie das einjährige Deutschlandstipendium ihrer Hochschule zur Förderung ihrer künstlerischen Tätigkeiten. Im Herbst gewann ihr Trio den HfM-Jazzpreis 2013, verliehen von der Hochschule für Musik Saar. Das Debütalbum des **Julia Kadel Trios** »Im Vertrauen« erschien am 29. August 2014 bei Blue Note/Universal. Damit nominierten sich **Kadel** und ihr Klaviertrio für den deutschen Echo Jazz 2015 in zwei Kategorien – **Kadel** als “Instrumentalist des Jahres national Piano” und das Trio als “Newcomer des Jahres”. 2016 erschien das zweite Album des Trios »Über und Unter«, ebenfalls bei Blue Note unter dem Dach von Universal Music.

Seit 2016 startete sie ihr Soloprojekt, mit dem sie bereits auf zahlreichen Festivals zu hören war. Von 2016 bis 2017 arbeitete sie in ihrem Duoprojekt mit dem türkisch-französischen Cellisten Anil Eraslan, mit dem sie bereits das Debütalbum »Noise Cloud« (Strasbourg, 2016) aufnahm. Seit 2015 wirkt **Kadel** ebenfalls im neu gegründeten Quartett der Free Jazz-Legende Günter Baby Sommer zusammen mit Friedhelm Schönfeld (sax/clar) und Walburga Walde (voc) mit. Darin begegnen sich zwei Generationen an Musiker*innen, nämlich solche "Vor der Mauer - nach der Mauer". 2017 erhielt sie das Künstlerstipendium des Australian Music Centre "AMPlify Germany" und entwickelte zusammen mit dem australischen Konzeptkünstler Julian Day das Experiment »How To Fail (Together)«. 2018 gründete sie ihr Duo mit dem deutschen Komponisten und Cellisten Thilo Thomas Krigar, mit dem sie sich in die Bereiche der improvisierten sowie Neuen Musik hineinbewegt.

Im September 2019 veröffentlichte ihr Trio sein drittes Album »Kaskaden« beim deutschen Traditions-Jazzlabel MPS in dessen legendärem, analogen Tonstudio im Schwarzwald. Damit nahm die Band der Berliner Pianistin als erster Act seit 35 Jahren wieder im historischen MPS Studio auf, worüber sogar der Spiegel Online berichtete. Aber auch als Solistin hat **Kadel** eine neue Dimension an Intensität und Komplexität in ihrer spezifischen musikalischen Ausdrucksweise gefunden. Sie spielte Konzerte in Philharmonien, wie der Elbphilharmonie Hamburg oder der Philharmonie Essen, wurde 2019 vom Magazin Jazzthing als eine der deutschen „Top Ten Key Players“ genannt. Auch über die deutschen Grenzen hinaus spielte sie international Konzerte und Festivals in Ländern wie Norwegen, England, Frankreich, Italien, Österreich, Ungarn, Tschechische Republik, Russland, Litauen und Türkei.
Seit Herbst 2020 spielt **Kadel** mit dem englischen Kontrabassisten Phil Donkin und dem amerikanischen Schlagzeuger Devin Gray in ihrem Trio zusammen. Im gleichen Jahr wurde die Pianistin und Komponistin für ihr innovatives musikalisches Schaffen mit dem Essener Preis "Jazz Pott 2020" ausgezeichnet.
 
**Grants**
Deutschlandstipendium 2013, Winner of HfM-Jazzpreis 2013, 2 Nominations for German Jazz Echo 2015, Solo Residency at Muziek Biennale Niederrhein (2016), Tour support from Goethe Institute Budapest (2016), Tour support from Goethe Institute Paris (2017), Artistic Development Grant from the Australian Music Centre "AMPlify Germany" (2017), Initiative Musik (2019) one year scholarship, Jazz Pott (2020)`,
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
						Bytes:       0,
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
