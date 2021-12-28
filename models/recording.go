package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Recording struct
type Recording struct {
	ID        uuid.UUID          `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"dffe2c0e-3713-4399-8ee2-279becbbb06e"`
	CreatedAt time.Time          `json:"created_at" example:"2020-12-10T19:00:00.000000+01:00"`
	UpdatedAt time.Time          `json:"updated_at" example:"2020-12-10T22:00:00.000000+01:00"`
	Channel   Channel            `json:"channel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ChannelID uuid.UUID          `json:"-" gorm:"type:uuid;unique_index:idx_recording_channel"`
	Formats   []*RecordingFormat `json:"formats,omitempty" gorm:"constraint:OnDelete:CASCADE" swaggerignore:"true"`
	// attributes
	CommonName string           `json:"common_name" gorm:"unique_index:idx_recording_channel" example:"2020-12-polizeigewalt"`
	Poster     string           `json:"poster" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/728edaf7-9ad9-f972-4d09-ba5940cd43f9/poster.png"`
	Preview    string           `json:"preview" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/728edaf7-9ad9-f972-4d09-ba5940cd43f9/preview.webp"`
	Duration   time.Duration    `json:"duration" swaggertype:"primitive,integer"`
	Public     bool             `json:"public"`
	Listed     bool             `json:"listed"`
	Viewers    uint64           `json:"viewers"`
	Lang       *RecordingLang   `json:"lang" gorm:"constraint:OnDelete:CASCADE"`
	Langs      []*RecordingLang `json:"-"`
	LangShorts []string         `json:"lang_shorts,omitempty" gorm:"-"`
	EventID    *uuid.UUID       `json:"-" gorm:"type:uuid"`
	Event      *Event           `json:"event"`
	Tags       []*Tag           `json:"tags" gorm:"many2many:recording_tags;constraint:OnDelete:CASCADE"`
	Speakers   []*Speaker       `json:"speakers" gorm:"many2many:recording_speakers;constraint:OnDelete:CASCADE"`
}

func (r *Recording) AfterFind(tx *gorm.DB) (err error) {
	count := len(r.Langs)
	if count > 0 {
		r.LangShorts = make([]string, count)
		for i, l := range r.Langs {
			r.LangShorts[i] = l.Lang
		}
	}
	return
}

// HasPermission - has user permission on stream
func (Recording) HasPermission(tx *gorm.DB, userID, objID uuid.UUID) (interface{}, error) {
	s := Recording{}
	count := 0
	if err := tx.Raw(`SELECT
		count(*)
		FROM user_channels uc
		INNER JOIN recordings s ON uc.channel_id=s.channel_id AND s.id=?
		WHERE uc.user_id = ?`,
		objID, userID).Scan(&count).Error; err != nil {
		return nil, err
	}
	if count != 1 {
		return nil, nil
	}
	if err := tx.First(&s, objID).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

// RecordingFormat struct - for format
type RecordingFormat struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"3a4f9157-65bf-4d15-a82b-1cd9295d07e0"`
	Recording   Recording `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RecordingID uuid.UUID `json:"-" gorm:"type:uuid"`
	// attributes
	Lang  string `json:"lang" gorm:"unique_index:idx_recording_lang" example:"de"`
	URL   string `json:"url" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/728edaf7-9ad9-f972-4d09-ba5940cd43f9/video_best.mp4"`
	Bytes int64  `json:"bytes"`
	//Quality    0: origin quality, 100: 4k, 140: 1440p, 150: 1080p, 160: 720p, 180: 480p. 200: 360p, 250: 240p, 300: 144p
	Quality    int64  `json:"quality" example:"0"`
	IsVideo    bool   `json:"is_video" example:"true"`
	Resolution string `json:"resolution" example:"1920x1080"`
}

// HasPermission - has user permission on stream
func (RecordingFormat) HasPermission(tx *gorm.DB, userID, objID uuid.UUID) (interface{}, error) {
	sl := RecordingFormat{}
	count := 0
	if err := tx.Raw(`SELECT
		count(*)
		FROM user_channels uc
		INNER JOIN recordings s ON uc.channel_id=s.channel_id
		INNER JOIN recording_formats sl ON s.id=sl.recording_id AND sl.id=?
		WHERE uc.user_id = ?`,
		objID, userID).Scan(&count).Error; err != nil {
		return nil, err
	}
	if count != 1 {
		return nil, nil
	}
	if err := tx.First(&sl, objID).Error; err != nil {
		return nil, err
	}
	return &sl, nil
}

// RecordingLang struct - for i18n data of a
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

// HasPermission - has user permission on stream
func (RecordingLang) HasPermission(tx *gorm.DB, userID, objID uuid.UUID) (interface{}, error) {
	sl := RecordingLang{}
	count := 0
	if err := tx.Raw(`SELECT
		count(*)
		FROM user_channels uc
		INNER JOIN recordings s ON uc.channel_id=s.channel_id
		INNER JOIN recording_langs sl ON s.id=sl.recording_id AND sl.id=?
		WHERE uc.user_id = ?`,
		objID, userID).Scan(&count).Error; err != nil {
		return nil, err
	}
	if count != 1 {
		return nil, nil
	}
	if err := tx.First(&sl, objID).Error; err != nil {
		return nil, err
	}
	return &sl, nil
}

func init() {
	migrations = append(migrations, []*gormigrate.Migration{
		{
			ID: "01-schema-0020-02-recording",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Recording{},
					&RecordingLang{},
					&RecordingFormat{})
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropTable("recording_formats"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("recording_langs"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("recording_speakers"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("recording_tags"); err != nil {
					return err
				}
				return tx.Migrator().DropTable("recordings")
			},
		},
	}...)
}
