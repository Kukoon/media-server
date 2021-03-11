package docs

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Kukoon/media-server/docs"
	"github.com/Kukoon/media-server/web"
)

func init() {
	web.ModuleRegister(func(r *gin.Engine, ws *web.Service) {
		r.GET("/api/help/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	})
}
