package models

import (
	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Channel struct
type Channel struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"df1555f5-7046-4f7a-adcc-195b73949723"`
	// attributes
	CommonName string       `json:"common_name" gorm:"unique" example:"kukoon"`
	Title      string       `json:"title" example:"Im Kukoon"`
	Logo       string       `json:"logo" example:"https://media.kukoon.de/static/css/kukoon/logo.png"`
	Recordings []*Recording `json:"recordings,omitempty" swaggerignore:"true"`
	Owners     []User       `json:"-" gorm:"many2many:user_channels;"`
}

// HasPermission - has user permission on channel
func (Channel) HasPermission(tx *gorm.DB, userID, objID uuid.UUID) (interface{}, error) {
	c := Channel{}
	count := 0
	if err := tx.Debug().Raw("SELECT count(*) FROM user_channels uc WHERE uc.user_id = ? AND uc.channel_id = ?", userID, objID).Scan(&count).Error; err != nil {
		return nil, err
	}
	if count != 1 {
		return nil, nil
	}
	if err := tx.First(&c, objID).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

var (
	// TestChannelID1 - uuid of channel 1 for testing
	TestChannelID1 = uuid.MustParse("df1555f5-7046-4f7a-adcc-195b73949723")

	// TestChannelID2 - uuid of channel 2 for testing
	TestChannelID2 = uuid.MustParse("c4eba06b-1ab3-4367-93e1-da584b96fcc8")
)

func init() {
	migrations = append(migrations, []*gormigrate.Migration{
		{
			ID: "01-schema-0010-01-channel",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Channel{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("channels")
			},
		},
	}...)

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0010-01-channel",
			Migrate: func(tx *gorm.DB) error {
				return tx.Create(&Channel{
					ID:         TestChannelID1,
					CommonName: "kukoon",
					Title:      "Im Kukoon",
					Logo:       "https://media.kukoon.de/static/css/kukoon/logo.png",
					Owners: []User{
						{ID: TestUserID1},
					},
				}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Exec("DELETE FROM user_channels WHERE channel_id = ?", TestChannelID1).Error; err != nil {
					return err
				}
				return tx.Delete(&Channel{
					ID: TestChannelID1,
				}).Error
			},
		},
		{
			ID: "10-data-0010-02-channel",
			Migrate: func(tx *gorm.DB) error {
				return tx.Create(&Channel{
					ID:         TestChannelID2,
					CommonName: "c3woc",
					Title:      "C3 Waffel Operation Center",
					Logo:       "https://c3woc.de/images/logo.svg",
					Owners: []User{
						{ID: TestUserID1},
					},
				}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Exec("DELETE FROM user_channels WHERE channel_id = ?", TestChannelID2).Error; err != nil {
					return err
				}
				return tx.Delete(&Channel{
					ID: TestChannelID2,
				}).Error
			},
		},
	}...)
}
