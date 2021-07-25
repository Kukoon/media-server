package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	testdataID := uuid.MustParse("27efbfff-d66c-c935-b308-9b1ee2bf78c8")
	testdataIDLang1 := uuid.MustParse("4f5ad673-2496-429a-a74f-0b48acdb807b")
	testdataIDFormat1 := uuid.MustParse("357af110-9481-4d0e-9fea-f61b30ee26f4")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0020-01-recording-5",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataID,
					ChannelID:  TestChannelID1,
					CommonName: "2021-02-geschichte_wird_gemacht",
					Poster:     "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					CreatedAt:  time.Date(2021, 2, 26, 19, 0, 0, 0, loc),
					Duration:   time.Hour + 12*time.Minute + 33*time.Second,
					Public:     true,
					Listed:     true,
					Tags: []*Tag{
						{ID: TestTagDiskussionID},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
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
				if err := tx.Create(&[]*RecordingFormat{
					{
						ID:          testdataIDFormat1,
						RecordingID: testdataID,
						Lang:        "de",
						Quality:     0,
						IsVideo:     true,
						URL:         "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/video_best.mp4",
						Bytes:       862470450,
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
