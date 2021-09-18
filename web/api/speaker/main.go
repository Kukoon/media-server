package speaker

import (
	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/gin-gonic/gin"
)

// Bind to webservice
func Bind(r *gin.Engine, ws *web.Service) {
	apiChannelListMy(r, ws)
	apiList(r, ws)
	apiDelete(r, ws)
}
