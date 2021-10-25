package stream

import (
	"time"

	"github.com/Kukoon/media-server/models"
	"github.com/google/uuid"
)

// Stream struct
type Stream struct {
	ID       uuid.UUID `json:"id" example:"dffe2c0e-3713-4399-8ee2-279becbbb06e"`
	ListenAt time.Time `json:"listen_at" example:"2020-12-10T00:00:00.000000+01:00"`
	StartAt  time.Time `json:"start_at" example:"2020-12-10T19:00:00.000000+01:00"`
	EndAt    time.Time `json:"end_at" example:"2020-12-10T22:00:00.000000+01:00"`
	Chat     bool      `json:"chat"`
	Running  bool      `json:"running"`
	// attributes
	CommonName string      `json:"common_name" example:"2020-12-polizeigewalt"`
	Poster     string      `json:"poster" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/728edaf7-9ad9-f972-4d09-ba5940cd43f9/poster.png"`
	Preview    string      `json:"preview" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/728edaf7-9ad9-f972-4d09-ba5940cd43f9/preview.webp"`
	EventID    *uuid.UUID  `json:"event_id" gorm:"type:uuid"`
	Tags       []uuid.UUID `json:"tags,omitempty" example:"df1555f5-7046-4f7a-adcc-195b73949723,0bca0cf4-a9b9-46d7-821f-18c59c08fc1d"`
	Speakers   []uuid.UUID `json:"speakers,omitempty" example:"b778369e-d44f-4d15-bf79-a9e8faec022d,09120c00-acb6-4865-b1b5-6bf39963e47f"`
}

func (s *Stream) Model() *models.Stream {
	if s == nil {
		return nil
	}
	m := &models.Stream{
		ID:         s.ID,
		ListenAt:   s.ListenAt,
		StartAt:    s.StartAt,
		EndAt:      s.EndAt,
		Chat:       s.Chat,
		Running:    s.Running,
		CommonName: s.CommonName,
		Poster:     s.Poster,
		Preview:    s.Preview,
		EventID:    s.EventID,
	}
	if s.Tags != nil {
		m.Tags = make([]*models.Tag, len(s.Tags))
		for i, id := range s.Tags {
			m.Tags[i] = &models.Tag{ID: id}
		}
	} else {
		m.Tags = []*models.Tag{}
	}
	if s.Speakers != nil {
		m.Speakers = make([]*models.Speaker, len(s.Speakers))
		for i, id := range s.Speakers {
			m.Speakers[i] = &models.Speaker{ID: id}
		}
	} else {
		m.Speakers = []*models.Speaker{}
	}
	return m
}
