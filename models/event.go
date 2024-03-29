package models

import (
	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Event struct
type Event struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	OwnerID     uuid.UUID `json:"-" gorm:"type:uuid"` // channel
	Owner       Channel   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name        string    `json:"name"`
	Logo        string    `json:"logo,omitempty"`
	URL         string    `json:"url,omitempty"`
	Description string    `json:"description,omitempty"`
}

// HasPermission - has user permission on event
func (Event) HasPermission(tx *gorm.DB, userID, objID uuid.UUID) (interface{}, error) {
	e := Event{}
	count := 0
	if err := tx.Raw(`SELECT
		count(*)
		FROM user_channels uc
		INNER JOIN events e ON uc.channel_id=e.owner_id AND e.id=?
		WHERE uc.user_id = ?`,
		objID, userID).Scan(&count).Error; err != nil {
		return nil, err
	}
	if count != 1 {
		return nil, nil
	}
	if err := tx.First(&e, objID).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

var (

	// TestEventID1 - uuid of event 1 for testing
	TestEventID1 = uuid.MustParse("bff61adc-76d5-4beb-aab0-3ef11b337204")

	// TestEventID2 - uuid of event 2 for testing
	TestEventID2 = uuid.MustParse("4abb3a05-60a3-4be5-a6aa-323b9755e0b5")
)

func init() {
	migrations = append(migrations, []*gormigrate.Migration{
		{

			ID: "01-schema-0017-02-event",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Event{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("events")
			},
		},
	}...)

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0010-01-event-01",
			Migrate: func(tx *gorm.DB) error {
				return tx.Create(&Event{
					ID:          TestEventID1,
					OwnerID:     TestChannelID1,
					Name:        "OUT LOUD",
					Description: `OUT LOUD ist eine Veranstaltungsreihe des Bremer Literaturkontors und wird gefördert durch den Senator für Kultur der Freien Hansestadt Bremen, die VGH-Stiftung, die Waldemar-Koch-Stiftung, die Karin und Uwe Hollweg Stiftung, unterstützt vom Literaturhaus Bremen und präsentiert von Bremen Zwei.`,
				}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Delete(&Event{
					ID: TestEventID1,
				}).Error
			},
		},
		{
			ID: "10-data-0010-01-event-02",
			Migrate: func(tx *gorm.DB) error {
				return tx.Create(&Event{
					ID:      TestEventID2,
					OwnerID: TestChannelID1,
					Name:    "Grand Piano Festival 2021",
					Description: `Es kommt etwas Großartiges auf uns zu! Vom 21. bis 25. Juli veranstalten wir in der Neustadt ein Konzertflügel-Festival.
						Damit das neben den fulminanten Künstler*innen auch noch für alle zugänglich sein kann, wird das Grand Piano ohne Eintritt und ohne Anmeldung im Neustadtspark stattfinden.`,
				}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Delete(&Event{
					ID: TestEventID2,
				}).Error
			},
		},
		{
			ID: "10-data-0010-02-event-01-02",
			Migrate: func(tx *gorm.DB) error {
				return tx.Model(&Event{}).
					Where("id in ?", []uuid.UUID{TestEventID1, TestEventID2}).
					Update("owner_id", TestChannelID1).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return nil
			},
		},
	}...)
}
