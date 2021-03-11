package docs

import (
	"github.com/gin-gonic/gin"

	// gorometheus
	"github.com/chenjiandongx/ginprom"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// db metrics
	"gorm.io/plugin/prometheus"

	"github.com/Kukoon/media-server/web"
)

func init() {
	web.ModuleRegister(func(r *gin.Engine, ws *web.Service) {
		r.Use(ginprom.PromMiddleware(nil))

		if ws.DB != nil {
			ws.DB.Use(prometheus.New(prometheus.Config{
				RefreshInterval: 15,
			}))

		}

		r.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
	})
}
