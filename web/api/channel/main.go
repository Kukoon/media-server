package channel

import (
	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/Kukoon/media-server/oven"
	"github.com/gin-gonic/gin"
)

// Bind to webservice
func Bind(r *gin.Engine, ws *web.Service, oven *oven.Service) {
	apiRestreamList(r, ws, oven)
	apiRestreamDelete(r, ws, oven)
	apiRestreamAdd(r, ws, oven)
	apiList(r, ws)
	apiListMy(r, ws)
	apiGet(r, ws)
	apiPut(r, ws)
}
