package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Grand Piano: Pulsar Trio
	testdataID := uuid.MustParse("f4d8d0e2-47bf-412c-a5e1-80c0f85a9e4c")
	testdataIDLang1 := uuid.MustParse("619e61ea-a6de-42e3-b48f-6b5e52f8d920")
	testdataIDSpeaker1 := uuid.MustParse("09120c00-acb6-4865-b1b5-6bf39963e47f")
	testdataIDFormat1 := uuid.MustParse("dd842a1d-7963-4d8d-8b79-58c7b20bc69d")

	// see stream 10

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-recording-15",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Recording{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/preview.webp",
					CreatedAt: time.Date(2021, 7, 22, 20, 30, 0, 0, loc),
					Duration:  time.Hour + 37*time.Minute + 28*time.Second,
					Public:    true,
					Listed:    true,
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker1,
							Name:    "Pulsar Trio",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&RecordingLang{
					ID:          testdataIDLang1,
					RecordingID: testdataID,
					Lang:        "de",
					Title:       "Pulsar Trio",
					Subtitle:    "22. Juli 2021, 20:30",
					Short:       `Treibend, pulsierend, vital. Raffinierte Rhythmen, freie Improvisation und Stücke, deren Melodien das Zeug zum Ohrwurm haben `,
					Long:        `Treibend, pulsierend, vital. Raffinierte Rhythmen, freie Improvisation und Stücke, deren Melodien das Zeug zum Ohrwurm haben – mit scheinbar so gegensätzlichen Instrumenten wie Sitar, Piano und Drums lässt das **Pulsar Trio** nicht nur einen neuartigen Klangraum entstehen, sondern auch eine groovige Fusion aus freiem Jazzdenken und originären Worldbeats. »Trotz ihrer vielen Taktwechsel, Breaks und Generalpausen wirkt [die Musik] nie verkopft oder kompliziert, sondern bleibt stets nachvollziehbar und zugänglich, vielfach – dank eines ausgeprägten Gespürs für griffige Melodiebögen – sogar geradezu eingängig.« (Harry Schmidt) Die creole – Preisträger konnten die hohe Qualität ihrer musikalischen Entdeckungsreise auf einer Vielzahl von Konzerten unter Beweis stellen. So gastierten sie u.a. auf den Leverkusener Jazztagen, dem Fusion-Festival, dem Glastonbury Festival, dem Rudolstadt Festival sowie dem Jazzfestival Izmir. Mit ihrem 3. Album »Zoo of Songs«, das im April 2018 bei t3 records erschien, erschließen sich die drei Musiker noch einmal ganz neue musikalische Sphären – kontemplativ, elektronisch, rhythmisch listig, energiegeladen und in satten Klangfarben präsentiert sich das Trio herausragend gereift und mit diesem pointierten Freigeist, der es so unverwechselbar macht. Die »Zoo of Songs« – Record-Release-Tour lässt sie ab 2018 auf internationalen Konzertbühnen unterwegs sein.`,
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
						Bytes:       1616888151,
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
