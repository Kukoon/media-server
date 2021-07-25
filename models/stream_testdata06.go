package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	testdataID := uuid.MustParse("1742d9b6-c9c6-45fb-a3a3-4a3e7fac2987")
	testdataIDLang1 := uuid.MustParse("0b7136a6-4c51-49ac-99e9-27ef833169f6")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-stream-06",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					Chat:      true,
					Running:   true,
					StartAt:   time.Date(2021, 6, 24, 0, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 6, 24, 0, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					Tags: []*Tag{
						{ID: TestTagDiskussionID},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataIDLang1,
					StreamID: testdataID,
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
				return tx.Delete(&Stream{
					ID: testdataID,
				}).Error
			},
		},
	}...)
}
