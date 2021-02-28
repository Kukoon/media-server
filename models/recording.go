package models

import (
	"time"

	"github.com/google/uuid"
)

type Recording struct {
	ID        uuid.UUID          `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"dffe2c0e-3713-4399-8ee2-279becbbb06e"`
	CreatedAt time.Time          `json:"created_at" example:"2020-12-10T19:00:00.000000+01:00"`
	UpdatedAt time.Time          `json:"updated_at" example:"2020-12-10T22:00:00.000000+01:00"`
	Channel   Channel            `json:"channel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ChannelID uuid.UUID          `json:"-" gorm:"type:uuid;unique_index:idx_recording_channel"`
	Formats   []*RecordingFormat `json:"formats,omitempty" swaggerignore:"true"`
	// attributes
	CommonName string         `json:"common_name" gorm:"unique_index:idx_recording_channel" example:"2020-12-polizeigewalt"`
	Poster     string         `json:"poster" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee_20201211_165251mp4"`
	Preview    string         `json:"preview" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee_20201211_165251.gif"`
	Duration   time.Duration  `json:"duration"`
	Public     bool           `json:"public"`
	Listed     bool           `json:"listed"`
	Lang       *RecordingLang `json:"lang"`
	EventID    *uuid.UUID     `json:"-" gorm:"type:uuid"`
	Event      *Event         `json:"event"`
	Tags       []*Tag         `json:"tags" gorm:"many2many:recording_tags;"`
	Speakers   []*Speaker     `json:"speakers" gorm:"many2many:recording_speakers;"`
}

type RecordingFormat struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"3a4f9157-65bf-4d15-a82b-1cd9295d07e0"`
	Recording   Recording `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RecordingID uuid.UUID `json:"-" gorm:"type:uuid"`
	// attributes
	Lang  string `json:"lang" gorm:"unique_index:idx_recording_lang" example:"de"`
	URL   string `json:"url" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee_20201211_165251.mp4"`
	Bytes int64  `json:"bytes"`
	//Quality    0: origin quality, 100: 4k, 140: 1440p, 150: 1080p, 160: 720p, 180: 480p. 200: 360p, 250: 240p, 300: 144p
	Quality    int64  `json:"quality" example:"0"`
	IsVideo    bool   `json:"is_video" example:"true"`
	Resolution string `json:"resolution" example:"1920x1080"`
}

type RecordingLang struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"3a4f9157-65bf-4d15-a82b-1cd9295d07e0"`
	Recording   Recording `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RecordingID uuid.UUID `json:"-" gorm:"type:uuid;unique_index:idx_recording_lang"`
	// attributes
	Lang     string `json:"lang" gorm:"unique_index:idx_recording_lang" example:"de"`
	Title    string `json:"title" example:"Polizeigewalt – ein deutsches Problem"`
	Subtitle string `json:"subtitle" example:"Ein deutsches Problem Diskussionsveranstaltung mit Laila Abdul-Rahman, Greta (Grün-Weiße Hilfe Bremen) und Mathilda (Kampagne für Opfer rassistischer Polizeigewalt - KOP Bremen)"`
	Short    string `json:"short" example:"Nachdem Mord an George Floyd ist es zu großen Protesten in den Vereinigten Staaten gekommen. Auch in Deutschland sterben schwarze Menschen in Polizeigewahrsam. [...]"`
	Long     string `json:"long" example:"Nachdem Mord an George Floyd ist es zu großen Protesten in den Vereinigten Staaten gekommen. Auch in Deutschland sterben schwarze Menschen in Polizeigewahrsam.  Ihre Namen sind weitgehend unbekannt: William Tonou-Mbobda, Hussam Fadl, Rooble Warsame, Oury Jalloh, Yaya Diabi, Amed A., Aamir Ageeb, Achidi John, Laya-Alama Condé, Mohamed Idrissi – die Liste ließe sich fortsetzen."`
}
