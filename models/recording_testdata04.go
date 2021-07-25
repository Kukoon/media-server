package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	testdataID := uuid.MustParse("13a70ec7-6e74-5114-021f-4d7910752df1")
	testdataIDLang1 := uuid.MustParse("345a3743-42dd-4ee9-97ae-c3785bd4235e")
	testdataIDSpeaker1 := uuid.MustParse("f597b912-97b1-4e16-b431-054692a5d049")
	testdataIDFormat1 := uuid.MustParse("fae90633-a15d-4dfb-b017-dc8561df95c3")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-recording-04",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:         testdataID,
					ChannelID:  TestChannelID1,
					CommonName: "2021-02-der_berg_der_nackten_wahrheiten",
					Poster:     "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					CreatedAt:  time.Date(2021, 2, 11, 19, 0, 0, 0, loc),
					Duration:   29*time.Minute + 56*time.Second,
					Public:     true,
					Listed:     true,
					Tags: []*Tag{
						{ID: TestTagBuchvorstellungID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker1,
							Name:    "Jan Backmann",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
					Lang:        "de",
					Title:       "Der Berg der nackten Wahrheiten",
					Subtitle:    "Die Geschichte des legendären Monte Verità aus der Sicht einer Ziege erzählt",
					Short:       `Gusto verbringt sein Leben auf dem Monte Verità im Tessin. Doch nicht alles verläuft so sorgenlos, wie der Hippie-Vorläufer es sich in seiner Traumwelt vorstellt hatte: Das Geld der vegetarisch-kommunistischen FKK-Gemeinschaft wird langsam knapp und als Gusto auch noch eine Ziege aus dem Dorfe bei sich aufnimmt, wächst die Wut der Bewohner\*innen von Ascona auf die Aussteiger\*innen.  ...`,
					Long: `
Gusto verbringt sein Leben auf dem Monte Verità im Tessin. Doch nicht alles verläuft so sorgenlos, wie der Hippie-Vorläufer es sich in seiner Traumwelt vorstellt hatte: Das Geld der vegetarisch-kommunistischen FKK-Gemeinschaft wird langsam knapp und als Gusto auch noch eine Ziege aus dem Dorfe bei sich aufnimmt, wächst die Wut der Bewohner\*innen von Ascona auf die Aussteiger\*innen. Nichtsdestotrotz schmiedet Gusto einen irrwitzigen Plan, wie er seine geliebte Ziege weiterhin bei sich behalten kann. Nach seinem ersten, sehr erfolgreichen Comic *Mühsam, Anarchist in Anführungsstrichen*, veröffentlicht der Autor nun eine Erzählung, die zehn Jahre früher spielt, vor dem Hintergrund der Aktivitäten auf dem Monte Verità, dem Treffpunkt der ersten Aussteiger\*innen im 20. Jahrhundert. Auch dieses Mal zapft Bachmann historische Quellen an, um daraus eine pointierte und bissige politische Komödie zu machen. Im Mittelpunkt steht nun allerdings eine Ziege, die Ziege der Vegetarier\*innen. Eine Leseprobe findet sich [hier](https://www.editionmoderne.ch/buch/der-berg-der-nackten-wahrheiten/).

**Jan Bachmann**
Geboren 1986 in Basel, hat an der Deutschen Film- und Fernsehakademie in Berlin studiert. 2013 bis 2015 war er Mitglied in einem brandenburgischen FKK-Verein. Sein erster Comic *Mühsam, Anarchist in Anführungsstrichen* ist 2018 bei der Edition Moderne erschienen und wurde unter anderem für den Max und Moritz-Preis nominiert. Aktuell arbeitet er an einem Buch zum Exil von Kaiser Wilhelm II in Holland.
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
						Bytes:       2020856776,
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
