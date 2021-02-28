package models

import (
	"github.com/google/uuid"
)

type Channel struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"df1555f5-7046-4f7a-adcc-195b73949723"`
	// attributes
	CommonName string       `json:"common_name" gorm:"unique" example:"kukoon"`
	Title      string       `json:"title" example:"Im Kukoon"`
	Logo       string       `json:"logo" example:"https://media.kukoon.de/static/css/kukoon/logo.png"`
	Recordings []*Recording `json:"recordings,omitempty" swaggerignore:"true"`
}
