package stream

import (
	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/gin-gonic/gin"
)

// Bind to webservice
func Bind(r *gin.Engine, ws *web.Service) {
	// meta
	apiList(r, ws)
	apiGet(r, ws)
	apiPost(r, ws)
	apiPut(r, ws)
	apiDelete(r, ws)
	apiPostToRecording(r, ws)
	// of channel
	apiChannelGet(r, ws)
	apiChannelListMy(r, ws)
	// lang
	apiLangList(r, ws)
	apiLangPost(r, ws)
	apiLangPut(r, ws)
	apiLangDelete(r, ws)
}
