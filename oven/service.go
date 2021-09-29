package oven

import (
	"time"

	"dev.sum7.eu/genofire/golang-lib/file"
	"dev.sum7.eu/genofire/golang-lib/worker"
	ovenAPI "dev.sum7.eu/genofire/oven-exporter/api"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Service for oven related commands
type Service struct {
	log         *zap.Logger
	w           *worker.Worker
	Client      ovenAPI.Client    `toml:"client"`
	StreamCheck file.TOMLDuration `toml:"stream_check"`
	DB          *gorm.DB          `toml:"-"`
}

// Run start all related workers on oven service
func (s *Service) Run(log *zap.Logger) {
	s.log = log
	if s.w != nil {
		// TODO error handling
		return
	}
	if s.StreamCheck > 0 {
		s.w = worker.NewWorker(time.Duration(s.StreamCheck), s.checkRunning)
		s.w.Start()
	}
}
