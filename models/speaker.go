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
	Name         string    `json:"name" example:"Dr. Jane Buoy"`
	Organisation string    `json:"organisation,omitempty" example:"Kukoon"`
}

func init() {
	migrations = append(migrations, []*gormigrate.Migration{
		{
			ID: "01-schema-0010-01-speaker",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Speaker{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("speakers")
			},
		},
	}...)
}
