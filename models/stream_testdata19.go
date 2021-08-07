package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Brav_a
	testdataID := uuid.MustParse("0ad35867-851f-4475-a41a-a0b93246a287")
	testdataIDLang1 := uuid.MustParse("0115e718-2b7f-4d3e-a7f9-74f6a376bb52")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0020-01-stream-19",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 8, 21, 18, 30, 0, 0, loc),
					ListenAt:  time.Date(2021, 0, 0, 0, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					Tags: []*Tag{
						{ID: TestTagVortragID},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataIDLang1,
					StreamID: testdataID,
					Lang:     "de",
					Title:    "Fussball & Polyamorie",
					Subtitle: "Brav_a (Berlin)",
					Short:    `Mitten in Corona haben wir die neueste Ausgabe unseres queer-feministischen Teen-mags gedruckt und veröffentlicht. Leider war es uns bisher nicht möglich, eine richtige Release-Veranstaltung zu organisieren.`,
					Long: `Mitten in Corona haben wir die neueste Ausgabe unseres queer-feministischen Teen-mags gedruckt und veröffentlicht. Leider war es uns bisher nicht möglich, eine richtige Release-Veranstaltung zu organisieren. Deswegen sind wir umso froher, das jetzt in Bremen tun zu können. Die Schwerpunkt-Themen der letzten Ausgabe waren Fußball und Polyamorie. Bei Fußball denken wir an Empowerment für Mädchen, an die Popularität von Ballsportarten und an Fantum ohne Nationalismus und Lokalpatriotismus. Wir haben hierzu eine Reihe von spannenden Texten und Interviews zusammengebracht. Das zweite Thema der Ausgabe ist Polyamorie. Poly- im Gegensatz zu Monogamie, was heißt das eigentlich? Die Beiträge im Zine diskutieren Idee und Praktiken von Polyamorie; außerdem gibt es als Beilage ein ganzes Zine-im-Zine, das das Thema in Comicform vertieft. Schöne und weniger schöne poly-Erlebnisse werden in Texten und Comics verhandelt. Als Lesung halten wir einige der Texte aus dieser Ausgabe und dazu noch Texte aus den vergangenen drei Ausgaben. Dazu wird es ein kleines Quiz geben mit Musik, interaktive Aufgaben und viel Spaß!

Die **Brav_a** ist ein queer-feministisches D.I.Y. (Do It Yourself) Zine, das sich im Stil einer Teenie-Zeitschrift teils ernst, teils ironisch mit Themen wie Liebe, Sex, (Körper-, Beziehungs-, Hetero-, Homo-) Normativität, Freund*innenschaften, der queer- feministischen Szene und vielem mehr beschäftigt. Das erste Heft erschien im Juli 2012, nun suchen wir schon für die 14. Ausgabe eure Beiträge! Beiträge sind auf Englisch oder Deutsch, ernst oder witzig sein und wir versuchen, dass diese am besten sowohl von Teenager*innen als auch erfahrene (Queer-) Feminist*innen zu verstehen sind.`,
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
