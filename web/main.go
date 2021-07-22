package web

import (
	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/api/status"
	"dev.sum7.eu/genofire/golang-lib/web/metrics"
	oven "dev.sum7.eu/genofire/oven-exporter/api"
	"github.com/gin-gonic/gin"

	"github.com/Kukoon/media-server/web/api/channel"
	"github.com/Kukoon/media-server/web/api/recording"
	"github.com/Kukoon/media-server/web/api/stream"
	"github.com/Kukoon/media-server/web/podcast"
	wsStream "github.com/Kukoon/media-server/web/ws/stream"
)

// Bind to webservice
func Bind(oven *oven.Client, vhost, app string) web.ModuleRegisterFunc {
	wsStreamBinder := wsStream.Bind(oven, vhost, app)
	return func(r *gin.Engine, ws *web.Service) {
		status.Register(r, ws)
		metrics.Register(r, ws)

		channel.Bind(r, ws)
		recording.Bind(r, ws)
		stream.Bind(r, ws)

		podcast.Bind(r, ws)
		wsStreamBinder(r, ws)

	}
}
