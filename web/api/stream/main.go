package stream

import (
	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/gin-gonic/gin"
)

// Bind to webservice
func Bind(r *gin.Engine, ws *web.Service) {
	apiChannelGet(r, ws)
	apiList(r, ws)
	apiGet(r, ws)
	apiPost(r, ws)
}
