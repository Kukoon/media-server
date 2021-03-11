package web

import (
	"github.com/bdlm/log"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	modules []ModuleRegisterFunc
)

type ModuleRegisterFunc func(*gin.Engine, *Service)

func ModuleRegister(f ModuleRegisterFunc) {
	modules = append(modules, f)
}

// Bind all routing
// @title Mediathek API
// @version 1.0
// @description This is the first version of the OpenSource Mediathek of Tomorrow
// @termsOfService http://swagger.io/terms/
// -host v2.media.kukoon.de
// @BasePath /
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func (ws *Service) Bind(r *gin.Engine) {
	for _, f := range modules {
		f(r, ws)
	}

	log.Infof("loaded %d modules", len(modules))
	r.Use(static.Serve("/", static.LocalFile(ws.Webroot, false)))
}
