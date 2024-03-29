package docs

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"dev.sum7.eu/genofire/golang-lib/web"
)

// Bind to webservice
func Bind(r *gin.Engine, ws *web.Service) {
	r.GET("/api/help/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
