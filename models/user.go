package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"88078ec0-2135-445f-bf05-632701c77695"`
	Username   string     `json:"username" gorm:"unique" example:"kukoon"`
	Password   string     `json:"-" example:"super secret password"`
	ForgetCode *uuid.UUID `json:"-" gorm:"forget_code;type:uuid"`
	Channels   []Channel  `json:"-" gorm:"many2many:user_channels;"`
}
