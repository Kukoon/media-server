package models

import (
	"github.com/google/uuid"
)

type Channel struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"df1555f5-7046-4f7a-adcc-195b73949723"`
	// attributes
	CommonName string       `json:"common_name" example:"kukoon"`
	Title      string       `json:"title" example:"Im Kukoon"`
	Secret     uuid.UUID    `json:"secret" gorm:"type:uuid;default:gen_random_uuid()" example:"d78b12f8-6904-4b75-81ce-6b22d9fe76ff"`
	Recordings []*Recording `json:"recordings,omitempty" swaggerignore:"true"`
}
