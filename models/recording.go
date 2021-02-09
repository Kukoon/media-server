package models

import (
	"time"

	"github.com/google/uuid"
)

type Recording struct {
	ID      uuid.UUID          `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"dffe2c0e-3713-4399-8ee2-279becbbb06e"`
	Formats []*RecordingFormat `json:"formats,omitempty" swaggerignore:"true"`
	// attributes
	CommonName string        `json:"common_name" example:"2020-12-polizeigewalt"`
	Duration   time.Duration `json:"duration"`
	CreatedAt  time.Time     `json:"created_at"`
}

type RecordingFormat struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"3a4f9157-65bf-4d15-a82b-1cd9295d07e0"`
	Recording   Recording `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RecordingID uuid.UUID `json:"-" gorm:"type:uuid"`
	// attributes
	URL        string `json:"url" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee_20201211_165251.mp4"`
	IsVideo    bool   `json:"is_video" example:"true"`
	Resolution string `json:"resolution" example:"1920x1080"`
}
