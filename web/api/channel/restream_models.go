package channel

import (
	oven "dev.sum7.eu/genofire/oven-exporter/api"
)

// Restream data
type Restream struct {
	ID       string `json:"id" example:"d592d376-0d93-4c2a-8156-5ea872fa748a Blub"`
	Name     string `json:"name" example:"Blub"`
	Protocol string `json:"protocol" example:"rtmp"`
	URL      string `json:"url" example:"rtmp://a.rtmp.youtube.com/live2"`
	// extras
	State    string `json:"state" example:"ready"`
	Sequence int    `json:"sequence" example:"1"`
}

func RestreamFromOven(data *oven.ResponsePushData) *Restream {
	name := data.ID
	if len(name) > 37 {
		name = name[37:]
	}
	return &Restream{
		ID:       data.ID,
		Name:     name,
		Protocol: data.Protocol,
		URL:      data.URL,
		State:    data.State,
		Sequence: data.Sequence,
	}
}

// RestreamAdd
type RestreamAdd struct {
	Name      string `json:"name" example:"Blub"`
	Protocol  string `json:"protocol" example:"rtmp"`
	URL       string `json:"url" example:"rtmp://a.rtmp.youtube.com/live2"`
	StreamKey string `json:"streamKey" example:"SUPERSECRET"`
}
