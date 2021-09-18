package models

import (
	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Speaker struct
type Speaker struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"b778369e-d44f-4d15-bf79-a9e8faec022d"`
	OwnerID      uuid.UUID `json:"-" gorm:"type:uuid"` // channel
	Owner        Channel   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name         string    `json:"name" example:"Dr. Jane Buoy"`
	Organisation string    `json:"organisation,omitempty" example:"Kukoon"`
}

// HasPermission - has user permission on speaker
func (Speaker) HasPermission(tx *gorm.DB, userID, objID uuid.UUID) (interface{}, error) {
	s := Speaker{}
	count := 0
	if err := tx.Raw(`SELECT
		count(*)
		FROM user_channels uc
		INNER JOIN speakers s ON uc.channel_id=s.owner_id AND s.id=?
		WHERE uc.user_id = ?`,
		objID, userID).Scan(&count).Error; err != nil {
		return nil, err
	}
	if count != 1 {
		return nil, nil
	}
	if err := tx.First(&s, objID).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func init() {
	migrations = append(migrations, []*gormigrate.Migration{
		{
			ID: "01-schema-0010-02-speaker",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Speaker{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("speakers")
			},
		},
	}...)
}
