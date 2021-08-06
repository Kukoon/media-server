package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Jan Bachmann - Der Kaiser im Exil
	testdataID := uuid.MustParse("023538b2-7635-4040-b25c-5794136738f1")
	testdataIDLang1 := uuid.MustParse("e19cd411-6efb-4a2e-b24a-b4b9af140822")
	testdataIDSpeaker1 := uuid.MustParse("f597b912-97b1-4e16-b431-054692a5d049") // duplicated - see recording 04

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0020-01-stream-21",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 8, 20, 19, 30, 0, 0, loc),
					ListenAt:  time.Date(2021, 0, 0, 0, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.jpg",
					Tags: []*Tag{
						{ID: TestTagBuchvorstellungID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker1,
							Name:    "Jan Bachmann",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataIDLang1,
					StreamID: testdataID,
					Lang:     "de",
					Title:    "Der Kaiser im Exil",
					Subtitle: "Graphic Novel Präsentation",
					Short:    `Nach der Niederlage im Ersten Weltkrieg muss Kaiser Wilhelm untertauchen und findet Zuflucht bei einer befreundeten Adelsfamilie.`,
					Long: `Nach der Niederlage im Ersten Weltkrieg muss Kaiser Wilhelm untertauchen und findet Zuflucht bei einer befreundeten Adelsfamilie. Seiner üblichen kaiserlichen Tätigkeiten beraubt, holzt Wilhelm kurzerhand den Wald ab, der das Gut umgibt, derweil die Gastgeberin — mit Gespür für die historische Dimension des Geschehens — das Werk »Der Kaiser im Exil« verfasst. In »Der Kaiser im Exil« verwebt Jan Bachmann Zitate, historische Überlieferungen und frei Dazuerfundenes zu einer lustvollen und kritischen Parodie des Hochmuts der (ehemals) Mächtigen.

**Jan Bachmann** hat an der Deutschen Film- und Fernsehakademie in Berlin studiert. Sein erster Comic »Mühsam, Anarchist in Anführungsstrichen« ist 2018 bei der Edition Moderne erschienen und wurde unter anderem für den Max und Moritz-Preis nominiert.

Eine Veranstaltung von __associazione delle talpe__ in Kooperation mit der __Rosa-Luxemburg-Initiative – Die Rosa-Luxemburg-Stiftung in Bremen__.`,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Delete(&Stream{
					ID: testdataID,
				}).Error
			},
		},
	}...)
}
