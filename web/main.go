package web

import (
	"net/http"

	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
)

// Webservice to store Configuration and Webserver wide objects
// (like DB Connection)
type Webservice struct {
	Listen string
	db     *gorm.DB
}

// Run to startup all related web parts
// (e.g. configure the server, metrics, and finally bind routing)
func (ws *Webservice) Run() error {
	gin.EnableJsonDecoderDisallowUnknownFields()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// catch crashed
	r.Use(gin.Recovery())

	r.Use(gin.Logger())

	r.Use(ginprom.PromMiddleware(nil))
	if ws.db != nil {
		ws.db.Use(prometheus.New(prometheus.Config{
			RefreshInterval: 15,
		}))
	}
	ws.bind(r)
	return r.Run(ws.Listen)
}

// Bind all routing
func (ws *Webservice) bind(r *gin.Engine) {
	r.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "running",
		})
	})
}
