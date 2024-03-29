package oven

import (
	"time"

	"github.com/Kukoon/media-server/models"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (s *Service) checkRunning() {
	// check status of stream server
	resp, err := s.Client.RequestDefaultListStreams()
	if err != nil {
		s.log.Warn("status check for oven stream server", zap.Error(err))
		return
	}
	ids := make([]uuid.UUID, len(resp.Data))
	// case to uuid
	for i, idStr := range resp.Data {
		id, err := uuid.Parse(idStr)
		if err == nil {
			ids[i] = id
		}
	}

	now := time.Now()

	// update on database running
	if err := s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Stream{}).Where("running = true").Update("running", false).Error; err != nil {
			return err
		}
		if len(resp.Data) == 0 {
			return nil
		}
		for _, id := range ids {
			stream := &models.Stream{}
			if err := tx.Where("channel_id = ?", id).
				Where("start_at < ?", now).
				Where("end_at > ?", now).
				Order("start_at DESC").First(&stream).Error; err != nil {
				return err
			}
			if err := tx.Model(stream).Update("running", true).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		s.log.Warn("update on database failed", zap.Error(err))
	}
}
