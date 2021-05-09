package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Channel struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"df1555f5-7046-4f7a-adcc-195b73949723"`
	// attributes
	CommonName string       `json:"common_name" gorm:"unique" example:"kukoon"`
	Title      string       `json:"title" example:"Im Kukoon"`
	Logo       string       `json:"logo" example:"https://media.kukoon.de/static/css/kukoon/logo.png"`
	Recordings []*Recording `json:"recordings,omitempty" swaggerignore:"true"`
	Owners     []User       `json:"-" gorm:"many2many:user_channels;"`
}

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
