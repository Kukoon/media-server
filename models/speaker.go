package models

import (
	"github.com/google/uuid"
)

type Speaker struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"b778369e-d44f-4d15-bf79-a9e8faec022d"`
	OwnerID      uuid.UUID `json:"-" gorm:"type:uuid"` // channel
	Name         string    `json:"name" example:"Dr. Jane Buoy"`
	Organisation string    `json:"organisation,omitempty" example:"Kukoon"`
}
