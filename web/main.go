package web

import (
	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/api/status"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"dev.sum7.eu/genofire/golang-lib/web/metrics"
	"github.com/gin-gonic/gin"

	"github.com/Kukoon/media-server/oven"
	"github.com/Kukoon/media-server/web/api/channel"
	"github.com/Kukoon/media-server/web/api/recording"
	"github.com/Kukoon/media-server/web/api/stream"
	"github.com/Kukoon/media-server/web/podcast"
	wsStream "github.com/Kukoon/media-server/web/ws/stream"
)

// Bind to webservice
func Bind(oven *oven.Service) web.ModuleRegisterFunc {
	return func(r *gin.Engine, ws *web.Service) {
		status.Register(r, ws)
		metrics.Register(r, ws)
		auth.Register(r, ws)

		channel.Bind(r, ws)
		recording.Bind(r, ws)
		stream.Bind(r, ws)

		podcast.Bind(r, ws)
		wsStream.Bind(r, ws)

	}
}
