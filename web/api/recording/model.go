package recording

import (
	"time"

	"github.com/Kukoon/media-server/models"
	"github.com/google/uuid"
)

// Recording struct
type Recording struct {
	ID       uuid.UUID `json:"id" example:"dffe2c0e-3713-4399-8ee2-279becbbb06e"`
	// attributes
	CommonName string        `json:"common_name" example:"2020-12-polizeigewalt"`
	Poster     string        `json:"poster" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/728edaf7-9ad9-f972-4d09-ba5940cd43f9/poster.png"`
	Preview    string        `json:"preview" example:"https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/728edaf7-9ad9-f972-4d09-ba5940cd43f9/preview.webp"`
	Duration   time.Duration `json:"duration"`
	Public     bool          `json:"public"`
	Listed     bool          `json:"listed"`
	EventID    *uuid.UUID    `json:"event_id" gorm:"type:uuid"`
	Tags       []uuid.UUID   `json:"tags,omitempty" example:"df1555f5-7046-4f7a-adcc-195b73949723,0bca0cf4-a9b9-46d7-821f-18c59c08fc1d"`
	Speakers   []uuid.UUID   `json:"speakers,omitempty" example:"b778369e-d44f-4d15-bf79-a9e8faec022d,09120c00-acb6-4865-b1b5-6bf39963e47f"`
}

func (s *Recording) Model() *models.Recording {
	if s == nil {
		return nil
	}
	m := &models.Recording{
		ID:         s.ID,
		CommonName: s.CommonName,
		Poster:     s.Poster,
		Preview:    s.Preview,
		Duration:   s.Duration,
		Public:     s.Public,
		Listed:     s.Listed,
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
