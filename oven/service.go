package oven

import (
	"time"

	"dev.sum7.eu/genofire/golang-lib/file"
	"dev.sum7.eu/genofire/golang-lib/worker"
	ovenAPI "dev.sum7.eu/genofire/oven-exporter/api"
	"gorm.io/gorm"
)

// Service for oven related commands
type Service struct {
	Client      ovenAPI.Client    `toml:"client"`
	StreamCheck file.TOMLDuration `toml:"stream_check"`
	DB          *gorm.DB          `toml:"-"`
	w           *worker.Worker
}

// Run start all related workers on oven service
func (s *Service) Run() {
	if s.w != nil {
		// TODO error handling
		return
	}
	if s.StreamCheck > 0 {
		s.w = worker.NewWorker(time.Duration(s.StreamCheck), s.checkRunning)
		s.w.Start()
	}
}
