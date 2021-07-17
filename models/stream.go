package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Stream struct
type Stream struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"dffe2c0e-3713-4399-8ee2-279becbbb06e"`
	Channel   Channel   `json:"channel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ChannelID uuid.UUID `json:"-" gorm:"type:uuid;unique_index:idx_stream_channel"`
	Secret    string    `json:"secret" example:"717897bf-198c-4f1b-bb4f-5a25cb197107"`
	StartAt   time.Time `json:"start_at" example:"2020-12-10T18:30:00.000000+01:00"`
	ListenAt  time.Time `json:"listen_at" example:"2020-12-10T19:00:00.000000+01:00"`
	Chat      bool      `json:"chat"`
	Running   bool      `json:"running"`
	// attributes
	CommonName string      `json:"common_name" gorm:"unique_index:idx_stream_channel" example:"2020-12-polizeigewalt"`
	Poster     string      `json:"poster" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/728edaf7-9ad9-f972-4d09-ba5940cd43f9/poster.png"`
	Preview    string      `json:"preview" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/728edaf7-9ad9-f972-4d09-ba5940cd43f9/preview.webp"`
	Lang       *StreamLang `json:"lang"`
	EventID    *uuid.UUID  `json:"-" gorm:"type:uuid"`
	Event      *Event      `json:"event,omitempty"`
	Tags       []*Tag      `json:"tags,omitempty" gorm:"many2many:stream_tags"`
	Speakers   []*Speaker  `json:"speakers,omitempty" gorm:"many2many:stream_speakers"`
}

// TableName of stream struct
func (Stream) TableName() string {
	return "streams"
}

func (s *Stream) GetPublic() *PublicStream {
	return &PublicStream{
		ID:        s.ID,
		Channel:   s.Channel,
		ChannelID: s.ChannelID,
		Secret:    s.Secret,
		StartAt:   s.StartAt,
		ListenAt:  s.ListenAt,
		Chat:      s.Chat,
		Running:   s.Running,
		// attributes
		Poster:   s.Poster,
		Preview:  s.Preview,
		Lang:     s.Lang,
		EventID:  s.EventID,
		Event:    s.Event,
		Tags:     s.Tags,
		Speakers: s.Speakers,
	}
}

// PublicStream struct without secrets
// TODO maybe better in api model
type PublicStream struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"dffe2c0e-3713-4399-8ee2-279becbbb06e"`
	Channel   Channel   `json:"channel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ChannelID uuid.UUID `json:"-" gorm:"type:uuid;unique_index:idx_stream_channel"`
	Secret    string    `json:"-"`
	StartAt   time.Time `json:"-"`
	ListenAt  time.Time `json:"listen_at" example:"2020-12-10T19:00:00.000000+01:00"`
	Chat      bool      `json:"chat"`
	Running   bool      `json:"running"`
	// attributes
	Poster   string      `json:"poster" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee_20201211_165251mp4"`
	Preview  string      `json:"preview" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee_20201211_165251.gif"`
	Lang     *StreamLang `json:"lang" gorm:"foreignkey:StreamID"`
	EventID  *uuid.UUID  `json:"-" gorm:"type:uuid"`
	Event    *Event      `json:"event,omitempty"`
	Tags     []*Tag      `json:"tags,omitempty" gorm:"many2many:stream_tags"`
	Speakers []*Speaker  `json:"speakers,omitempty" gorm:"many2many:stream_speakers"`
}

// TableName of public stream struct (to fetch from same table then stream)
func (PublicStream) TableName() string {
	return "streams"
}

// StreamLang struct - for i18n data of a stream
type StreamLang struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"3a4f9157-65bf-4d15-a82b-1cd9295d07e0"`
	CreatedAt time.Time `json:"created_at" example:"2020-12-10T19:00:00.000000+01:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-12-10T22:00:00.000000+01:00"`
	Stream    Stream    `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	StreamID  uuid.UUID `json:"-" gorm:"type:uuid;unique_index:idx_stream_lang"`
	// attributes
	Lang     string `json:"lang" gorm:"unique_index:idx_stream_lang" example:"de"`
	Title    string `json:"title" example:"Pushbacks, Internierung, Zugangshürden"`
	Subtitle string `json:"subtitle" example:"Zum Stand des europäischen Migrations- und Grenzregimes"`
	Short    string `json:"short" example:"Nach dem katastrophalen Brand des Flüchtlingslagers Moria auf Lesbos hatte die Europäische Kommission erneut einen Neustart in der europäischen Migrations- und Asylpolitik versucht. [...]"`
	Long     string `json:"long" example:"Nach dem katastrophalen Brand des Flüchtlingslagers Moria auf Lesbos hatte die Europäische Kommission erneut einen Neustart in der europäischen Migrations- und Asylpolitik versucht."`
}

func init() {
	migrations = append(migrations, []*gormigrate.Migration{
		{
			ID: "01-schema-0030-01-stream",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Stream{}, &StreamLang{})
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropTable("stream_langs"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("stream_speakers"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("stream_tags"); err != nil {
					return err
				}
				return tx.Migrator().DropTable(Stream{}.TableName())
			},
		},
	}...)
}
