package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Bremen liest! 4. Bremer Literaturnacht
	testdataID := uuid.MustParse("9fa05bb3-e616-4249-bc77-aa102596fe21")
	testdataIDLang1 := uuid.MustParse("97afa470-c857-442c-8665-dea849f56c29")
	testdataIDSpeaker1 := uuid.MustParse("b1618955-6f9d-459e-ae93-b58cf5aeed84")
	testdataIDSpeaker2 := uuid.MustParse("f4170bac-415f-4c5d-8643-3c0c76492f65")
	testdataIDSpeaker3 := uuid.MustParse("bb75f586-738c-46bb-b9f9-922d5440dc74")
	testdataIDSpeaker4 := uuid.MustParse("c36f16e4-0882-4fde-bc1d-e74cb7961128")
	testdataIDSpeaker5 := uuid.MustParse("f901d519-1250-4152-903f-54103b186bdf")
	testdataIDSpeaker6 := uuid.MustParse("0fa16b1a-a777-401c-99c7-6cb7ef43e380")

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0020-01-stream-22",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataID,
					ChannelID: TestChannelID1,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 9, 3, 17, 55, 0, 0, loc),
					ListenAt:  time.Date(2021, 0, 0, 0, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/" + TestChannelID1.String() + "/" + testdataID.String() + "/poster.png",
					Tags: []*Tag{
						{ID: TestTagBuchvorstellungID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker1,
							Name:    "Farhan Hebbo",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker2,
							Name:    "Florian Reinartz",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker3,
							Name:    "Ursula Overhage",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker4,
							Name:    "Loubna Khaddaj",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker5,
							Name:    "Artur Becker",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataIDSpeaker6,
							Name:    "Helge Hommers",
							Organisation: "Literaturkontor",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataIDLang1,
					StreamID: testdataID,
					Lang:     "de",
					Title:    "Bremen liest! 4. Bremer Literaturnacht",
					Subtitle: "Bremer Autorinnen und Autoren lesen aus ihren Romanen, Kurzgeschichten, Kinder- und Sachbüchern.",
					Short:    `Bei der 4. Literaturnacht »Bremen liest!« lesen Autorinnen und Autoren über die ganze Stadt verteilt aus Romanen, Krimis, Kinderbüchern, Kurzgeschichten, Lyrik und Spoken Word Texten. In 20 Buchhandlungen und Bibliotheken sowie auf 5 Open-Air-Bühnen finden von 16 bis 23 Uhr verschiedenste Veranstaltungen statt.`,
					Long: `Bei der 4. Literaturnacht »Bremen liest!« lesen Autorinnen und Autoren über die ganze Stadt verteilt aus Romanen, Krimis, Kinderbüchern, Kurzgeschichten, Lyrik und Spoken Word Texten. In 20 Buchhandlungen und Bibliotheken sowie auf 5 Open-Air-Bühnen finden von 16 bis 23 Uhr verschiedenste Veranstaltungen statt. Das Bremer Literaturkontor lädt zu einer literarischen Tour durch die Stadt ein, um die Vielfalt der Bremer Literaturszene zu entdecken. Seien Sie dabei, wenn die ganze Hansestadt einen Abend lang im Zeichen des Vorlesens steht. 
 
Das gesamte Programm ist unter [www.bremenliest.de](https://www.bremenliest.de) zu finden im Parkcafe des Kukoon moderiert Helge Hommers (Literaturkontor):

**18:00 Uhr:** 
* Farhan Hebbo: »Einfache Geschichte« (Lyrik)
* Florian Reinartz: »Bianca, Susanne und der Dalai Lama« & weitere Kurzgeschichten

**19:30 Uhr:**
* Ursula Overhage: »Sie spielte wie im Rausch – Die Schauspielerin Maria Orska« (Sachbuch)
* Loubna Khaddaj: »Randvoll« aus: »Texte nach Hanau« (Prosa)

**21:00 Uhr:**
* Artur Becker: »Drang nach Osten« (Roman)`,
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
