package web

import (
	"github.com/bdlm/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// acme
	"github.com/gin-gonic/autotls"
	"golang.org/x/crypto/acme/autocert"
)

// Service to store Configuration and Webserver wide objects
// (like DB Connection)
type Service struct {
	// config
	Listen    string `toml:"listen"`
	AccessLog bool   `toml:"access_log"`
	Webroot   string `toml:"webroot"`
	ACME      struct {
		Enable  bool     `toml:"enable"`
		Domains []string `toml:"domains"`
		Cache   string   `toml:"cache"`
	} `toml:"acme"`
	Session struct {
		Name   string `toml:"name"`
		Secret string `toml:"secret"`
	} `toml:"session"`
	// internal
	DB *gorm.DB `toml:"-"`
}

// Run to startup all related web parts
// (e.g. configure the server, metrics, and finally bind routing)
func (config *Service) Run() error {
	gin.EnableJsonDecoderDisallowUnknownFields()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// catch crashed
	r.Use(gin.Recovery())

	if config.AccessLog {
		r.Use(gin.Logger())
		log.Debug("request logging enabled")
	}
	config.LoadSession(r)
	config.Bind(r)

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
