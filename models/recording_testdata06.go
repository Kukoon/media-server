package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	testdataID := uuid.MustParse("81b262e9-e010-1fa2-84a5-d8cee1a94835")
	testdataIDLang1 := uuid.MustParse("0ce4b366-9238-4aa4-a6d6-94227c1b0681")
	testdataIDSpeaker1 := uuid.MustParse("d8ba2b91-78f7-4bcd-9dc4-5af1d3c904a9")
	testdataIDSpeaker2 := uuid.MustParse("62d9ce45-1465-40f8-bf99-22607e7be91d")
	testdataIDFormat1 := uuid.MustParse("449e3361-f2e2-44ee-a5d7-3c013cfe1fdc")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-recording-06",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataID,
					ChannelID:  TestChannelID1,
					CommonName: "2021-03-verschwoerungserzaehlung",
					Poster:     "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					CreatedAt:  time.Date(2021, 3, 3, 19, 0, 0, 0, loc),
					Duration:   time.Hour + 14*time.Minute + 17*time.Second,
					Public:     true,
					Listed:     true,
					Tags: []*Tag{
						{ID: TestTagVortragID},
						{ID: TestTagDiskussionID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker1,
							Name:    "Johanna Bröse",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker2,
							Name:    "Andrea Strübe",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
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
				if err := tx.Create(&[]*RecordingFormat{
					{
						ID:          testdataIDFormat1,
						RecordingID: testdataID,
						Lang:        "de",
						Quality:     0,
						IsVideo:     true,
						URL:         "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/video_best.mp4",
						Bytes:       1426234816,
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
