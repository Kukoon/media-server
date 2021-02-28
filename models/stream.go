package models

import (
	"time"

	"github.com/google/uuid"
)

type Stream struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"dffe2c0e-3713-4399-8ee2-279becbbb06e"`
	Channel   Channel   `json:"channel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ChannelID uuid.UUID `json:"-" gorm:"type:uuid;unique_index:idx_recording_channel"`
	Secret    string    `json:"secret" example:"717897bf-198c-4f1b-bb4f-5a25cb197107"`
	StartAt   time.Time `json:"start_at" example:"2020-12-10T18:30:00.000000+01:00"`
	ListenAt  time.Time `json:"listen_at" example:"2020-12-10T19:00:00.000000+01:00"`
	Running   bool      `json:"running"`
	// attributes
	CommonName string      `json:"common_name" gorm:"unique_index:idx_recording_channel" example:"2020-12-polizeigewalt"`
	Poster     string      `json:"poster" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee_20201211_165251mp4"`
	Preview    string      `json:"preview" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee_20201211_165251.gif"`
	Lang       *StreamLang `json:"lang"`
	EventID    *uuid.UUID  `json:"-" gorm:"type:uuid"`
	Event      *Event      `json:"event,omitempty"`
	Tags       []*Tag      `json:"tags,omitempty" gorm:"many2many:stream_tags"`
	Speakers   []*Speaker  `json:"speakers,omitempty" gorm:"many2many:stream_speakers"`
}

func (s *Stream) GetPublic() *PublicStream {
	return &PublicStream{
		ID:        s.ID,
		Channel:   s.Channel,
		ChannelID: s.ChannelID,
		Secret:    s.Secret,
		StartAt:   s.StartAt,
		ListenAt:  s.ListenAt,
		Running:   s.Running,
		// attributes
		CommonName: s.CommonName,
		Poster:     s.Poster,
		Preview:    s.Preview,
		Lang:       s.Lang,
		EventID:    s.EventID,
		Event:      s.Event,
		Tags:       s.Tags,
		Speakers:   s.Speakers,
	}
}

type PublicStream struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"dffe2c0e-3713-4399-8ee2-279becbbb06e"`
	Channel   Channel   `json:"channel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ChannelID uuid.UUID `json:"-" gorm:"type:uuid;unique_index:idx_recording_channel"`
	Secret    string    `json:"-"`
	StartAt   time.Time `json:"-"`
	ListenAt  time.Time `json:"listen_at" example:"2020-12-10T19:00:00.000000+01:00"`
	Running   bool      `json:"running"`
	// attributes
	CommonName string      `json:"common_name" gorm:"unique_index:idx_recording_channel" example:"2020-12-polizeigewalt"`
	Poster     string      `json:"poster" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee_20201211_165251mp4"`
	Preview    string      `json:"preview" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee_20201211_165251.gif"`
	Lang       *StreamLang `json:"lang"`
	EventID    *uuid.UUID  `json:"-" gorm:"type:uuid"`
	Event      *Event      `json:"event,omitempty"`
	Tags       []*Tag      `json:"tags,omitempty" gorm:"many2many:stream_tags"`
	Speakers   []*Speaker  `json:"speakers,omitempty" gorm:"many2many:stream_speakers"`
}

type StreamLang struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"3a4f9157-65bf-4d15-a82b-1cd9295d07e0"`
	CreatedAt time.Time `json:"created_at" example:"2020-12-10T19:00:00.000000+01:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-12-10T22:00:00.000000+01:00"`
	Stream    Stream    `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	StreamID  uuid.UUID `json:"-" gorm:"type:uuid;unique_index:idx_recording_lang"`
	// attributes
	Lang     string `json:"lang" gorm:"unique_index:idx_recording_lang" example:"de"`
	Title    string `json:"title" example:"Pushbacks, Internierung, Zugangshürden"`
	Subtitle string `json:"subtitle" example:"Zum Stand des europäischen Migrations- und Grenzregimes"`
	Short    string `json:"short" example:"Nach dem katastrophalen Brand des Flüchtlingslagers Moria auf Lesbos hatte die Europäische Kommission erneut einen Neustart in der europäischen Migrations- und Asylpolitik versucht. [...]"`
	Long     string `json:"Long" example:"Nach dem katastrophalen Brand des Flüchtlingslagers Moria auf Lesbos hatte die Europäische Kommission erneut einen Neustart in der europäischen Migrations- und Asylpolitik versucht."`
}
