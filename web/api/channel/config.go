package channel

import (
	"strings"

	"github.com/Kukoon/media-server/models"
)

type StreamIngress struct {
	WebRTC string `json:"webrtc"`
	RTMP   string `json:"rtmp"`
}

// ForChannel fills StreamIngress by replacer with values for a defined channel
func (si *StreamIngress) ForChannel(r *strings.Replacer) StreamIngress {
	return StreamIngress{
		WebRTC: r.Replace(si.WebRTC),
		RTMP:   r.Replace(si.RTMP),
	}
}

type StreamSource struct {
	Label     string `json:"label"`
	Type      string `json:"type"`
	File      string `json:"file"`
	FrameRate int    `json:"framerate,omitempty"`
}

// ForChannel fills StreamSource by replacer with values for a defined channel
func (ss *StreamSource) ForChannel(r *strings.Replacer) StreamSource {
	return StreamSource{
		Label:     ss.Label,
		Type:      ss.Type,
		File:      r.Replace(ss.File),
		FrameRate: ss.FrameRate,
	}
}

type ConfigStream struct {
	Streams []StreamSource `config:"stream_sources"`
	Ingress StreamIngress  `config:"ingress"`
}

// ForChannel fills ConfigStream by placeholder with values for a defined channel
func (cs *ConfigStream) ForChannel(obj models.Channel) ConfigStream {
	r := strings.NewReplacer("{ID}", obj.ID.String())
	channelStream := ConfigStream{
		Ingress: cs.Ingress.ForChannel(r),
	}
	for _, s := range cs.Streams {
		channelStream.Streams = append(channelStream.Streams, s.ForChannel(r))
	}
	return channelStream
}
