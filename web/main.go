package web

import (
	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/api/status"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"dev.sum7.eu/genofire/golang-lib/web/metrics"
	"github.com/gin-gonic/gin"

	"github.com/Kukoon/media-server/oven"
	"github.com/Kukoon/media-server/web/api/channel"
	"github.com/Kukoon/media-server/web/api/event"
	"github.com/Kukoon/media-server/web/api/recording"
	"github.com/Kukoon/media-server/web/api/speaker"
	"github.com/Kukoon/media-server/web/api/stream"
	"github.com/Kukoon/media-server/web/api/tag"
	"github.com/Kukoon/media-server/web/docs"
	"github.com/Kukoon/media-server/web/podcast"
	wsStream "github.com/Kukoon/media-server/web/ws/stream"
)

// Bind to webservice
func Bind(oven *oven.Service) web.ModuleRegisterFunc {
	return func(r *gin.Engine, ws *web.Service) {
		docs.Bind(r, ws)

		status.Register(r, ws)
		metrics.Register(r, ws)
		auth.Register(r, ws)

		channel.Bind(r, ws)
		recording.Bind(r, ws)
		stream.Bind(r, ws)

		speaker.Bind(r, ws)
		event.Bind(r, ws)
		tag.Bind(r, ws)

		podcast.Bind(r, ws)
		wsStream.Bind(r, ws)

	}
}
