package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// 2021-06-17\ Erinnern\ stören
	testdataID := uuid.MustParse("39cbaa42-b3dd-4679-8611-0a00fe365ad5")
	testdataIDLang1 := uuid.MustParse("d5c676a5-71a2-4020-aa34-85623bdea5ae")
	testdataIDSpeaker1 := uuid.MustParse("01c25f76-1d9e-43a2-b6e9-d869dca35622")
	testdataIDSpeaker2 := uuid.MustParse("3119cb20-ab67-43e3-be41-09b2a1b174ca")
	testdataIDFormat1 := uuid.MustParse("d6b02b1a-d2dd-4660-a888-d2b34a66552a")

	// no stream

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-recording-11",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataID,
					ChannelID:  TestChannelID1,
					CommonName: "2021-06_erinnern_stoeren",
					Poster:     "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					CreatedAt:  time.Date(2021, 6, 17, 19, 0, 0, 0, loc),
					Duration:   time.Hour + 32*time.Minute + 16*time.Second,
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
							Name:    "Lydia Lierke",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker2,
							Name:    "Dan Thy Nguyen",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
					Lang:        "de",
					Title:       "Erinnern stören",
					Subtitle:    "Der Mauerfall aus migrantischer und jüdischer Perspektive",
					Short:       `Der Mauerfall vor 30 Jahren bedeutete eine gewaltvolle Zäsur für migrantisches und jüdisches Leben in Ost und West. Während die einen vereinigt wurden, wurden die anderen ausgeschlossen.`,
					Long: `Der Mauerfall vor 30 Jahren bedeutete eine gewaltvolle Zäsur für migrantisches und jüdisches Leben in Ost und West. Während die einen vereinigt wurden, wurden die anderen ausgeschlossen. Das vorliegende Buch möchte ausgegrenzte Perspektiven auf die deutsch-deutsche Vereinigung wieder sichtbar machen und an die Kämpfe um Teilhabe in den 1980er Jahren, einschneidende Erlebnisse um die Wende und die Selbstbehauptung gegen den Rassismus der 1990er Jahre erinnern. So beinhaltet der Band Geschichten von Bürgerrechts- und Asylkämpfen ehemaliger Gastarbeiterinnen, von Geflüchteten in BRD und DDR, Beiträge über den Eigensinn von Vertragsarbeiterinnen, von damaligen internationalen Studierenden, über jüdisches Leben in Ost und West sowie über die Kämpfe von Sinti und Roma im geteilten Deutschland. Mit Beiträgen von Sharon Adler, Emmanuel Adu Agyeman, Pablo Dominguez Andersen, Felix Axster, Mathias Berek, Gabriel Berger, Róza Berger-Fiedler, Hamze Bytyçi, Leah Carola Czollek, Max Czollek, Nuray Demir, Dostluk Sineması, Gülriz Egilmez, Naika Foroutan, Mirna Funk, Elisa Gutsche, Kathleen Heft, Initiative 12. August, Anetta Kahane, Dmitrij Kapitelman, Kadriye Karcı, Andrea Caroline Keppler, Evrim Efsun Kızılay, Jana König, David Kowalski, Janko Lauenberger, Lydia Lierke, Jessica Massochua, Paulino Miguel, Dan Thy Nguyen, Hannah Peaceman, Massimo Perinelli, Patrice G. Poutrus, Sabuha, Elisabeth Steffen, Ceren Türkmen, Nea Weissberg, Alexandra Weltz-Rombach und Cynthia Zimmermann.
Mit Illustrationen von Nino Paula Bulling und Burcu Türker.`,
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
						Bytes:       2023362978,
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
