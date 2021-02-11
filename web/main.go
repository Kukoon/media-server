package web

import (
	"net/http"

	"github.com/bdlm/log"
	"github.com/gin-gonic/gin"
	// prometheus
	"github.com/chenjiandongx/ginprom"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// db metrics
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
	// acme
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/autotls"
	"golang.org/x/crypto/acme/autocert"
	// swagger
	_ "github.com/Kukoon/media-server/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Webservice to store Configuration and Webserver wide objects
// (like DB Connection)
type Webservice struct {
	// config
	Listen    string `toml:"listen"`
	AccessLog bool   `toml:"access_log"`
	Webroot   string `toml:"webroot"`
	ACME      struct {
		Enable  bool     `toml:"enable"`
		Domains []string `toml:"domains"`
		Cache   string   `toml:"cache"`
	} `toml:"acme"`
	// internal
	DB *gorm.DB `toml:"-"`
}

// Run to startup all related web parts
// (e.g. configure the server, metrics, and finally bind routing)
func (config *Webservice) Run() error {
	gin.EnableJsonDecoderDisallowUnknownFields()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// catch crashed
	r.Use(gin.Recovery())

	if config.AccessLog {
		r.Use(gin.Logger())
		log.Debug("request logging enabled")
	}

	r.Use(ginprom.PromMiddleware(nil))
	if config.DB != nil {
		config.DB.Use(prometheus.New(prometheus.Config{
			RefreshInterval: 15,
		}))
	}
	config.bind(r)

	if config.ACME.Enable {
		if config.Listen != "" {
			log.Panic("For ACME / Let's Encrypt it is not possible to set `listen`")
		}
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(config.ACME.Domains...),
			Cache:      autocert.DirCache(config.ACME.Cache),
		}
		return autotls.RunWithManager(r, &m)
	} else {
		return r.Run(config.Listen)
	}
}

// Bind all routing
// @title Mediathek API
// @version 1.0
// @description This is the first version of the OpenSource Mediathek of Tomorrow
// @termsOfService http://swagger.io/terms/
// -host v2.media.kukoon.de
// @BasePath /api
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func (ws *Webservice) bind(r *gin.Engine) {
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "running",
		})
	})
	r.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
	r.GET("/rss/:slug", ws.rssChannel)

	api := r.Group("/api")
	api.GET("/help/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := api.Group("/v1")
	{
		v1.GET("/channels", ws.apiChannelList)
		v1.GET("/channel/:slug", ws.apiChannelGet)
		v1.GET("/recordings", ws.apiRecordingList)
		v1.GET("/recording/:slug", ws.apiRecordingGet)
	}

	r.Use(static.Serve("/", static.LocalFile(ws.Webroot, false)))
}
