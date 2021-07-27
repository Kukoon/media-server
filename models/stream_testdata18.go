package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Gerhard Stapelfeldt: Revolte der Natur und konformistischer Protest – Über die Klimaschutzbewegung „Fridays for Future“
	testdataID := uuid.MustParse("aeea4b50-6974-4f95-adea-ab7f330f9509")
	testdataIDLang1 := uuid.MustParse("b2630449-2e99-49e9-9a87-b60b5e455dec")
	testdataIDSpeaker1 := uuid.MustParse("d185a32a-62d8-4a6c-8ca9-629b2203e89e")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0020-01-stream-18",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 8, 13, 19, 30, 0, 0, loc),
					ListenAt:  time.Date(2021, 8, 13, 0, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					Tags: []*Tag{
						{ID: TestTagVortragID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker1,
							Name:    "Gerhard Stapelfeldt",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataIDLang1,
					StreamID: testdataID,
					Lang:     "de",
					Title:    "Revolte der Natur und konformistischer Protest",
					Subtitle: "Über die Klimaschutzbewegung „Fridays for Future“",
					Short:    `Die Erkenntnis, dass die immer deutlicher spürbare „Klimakrise“anthropogen verursacht ist, wird gegenwärtig nur noch von Rechtspopulist*innen bestritten. Dennoch sind nicht die „anthropogenen“, also gesellschaftlichen und ökonomischen Ursachen gut erforscht, sondern nur die Erscheinungen und die naturwissenschaftlich zu bestimmenden Zusammenhänge.`,
					Long: `Die Erkenntnis, dass die immer deutlicher spürbare „Klimakrise“anthropogen verursacht ist, wird gegenwärtig nur noch von Rechtspopulist*innen bestritten. Dennoch sind nicht die „anthropogenen“, also gesellschaftlichen und ökonomischen Ursachen gut erforscht, sondern nur die Erscheinungen und die naturwissenschaftlich zu bestimmenden Zusammenhänge. Analog fordern die Sprecher*innen der Klimaschutzbewegung Fridays for Future auch nur, den Klimaforscher*innen solle endlich zugehört werden. Weil der Widerstand der Bewegung auf den Gebieten der Gesellschaftstheorie und Politischen Ökonomie analphabetisch ist und sich nicht gegen die gesellschaftlichen und ökonomischen Ursachen der Krise richtet, trifft er auf keinen nennenswerten gesellschaftlichen und politischen Widerstand. Das Aufbegehren der Klimaschutz-Aktivis*innen ist ein konformistischer Protest: eine neoliberale Kritik des globalisierten Neoliberalismus. Im Vortrag wird eine Metakritik dieses gesellschaftlichen und ökonomischen Analphabetismus versucht: Die gesellschaftliche Sprachlosigkeit wird nicht abstrakt denunziert, sondern aus den bestehenden neoliberalen Verhältnissen aufgeklärt, um jene Bewusstlosigkeit zu überwinden und die Kritik der Bewegung gesellschaftstheoretisch und politisch-praktisch zu radikalisieren.

**Gerhard Stapelfeldt** lehrte bis 2009 als Professor am Institut für Soziologie der Universität Hamburg. Seitdem arbeitet er als freier Schriftsteller in Hamburg.`,
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
