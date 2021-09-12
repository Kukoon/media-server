package recording

import (
	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"
)

func bindTest(r *gin.Engine, ws *web.Service) {
	Bind(r, ws)
	auth.Register(r, ws)
}
