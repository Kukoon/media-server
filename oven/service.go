package oven

import (
	ovenAPI "dev.sum7.eu/genofire/oven-exporter/api"
)

// Service for oven related commands
type Service struct {
	Client ovenAPI.Client `toml:"client"`
}
