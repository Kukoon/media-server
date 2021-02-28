package models

import (
	"github.com/google/uuid"
)

type Tag struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"df1555f5-7046-4f7a-adcc-195b73949723"`
	Lang *TagLang  `json:"lang"`
}
type TagLang struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"df1555f5-7046-4f7a-adcc-195b73949723"`
	Tag   Tag       `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TagID uuid.UUID `json:"-" gorm:"type:uuid;unique_index:idx_tag_lang"`
	// attributes
	Lang string `json:"lang" gorm:"unique_index:idx_tag_lang" example:"de"`
	Name string `json:"name" example:"Buchvorstellung"`
}
