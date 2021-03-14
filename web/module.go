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

func (ws *Service) Bind(r *gin.Engine) {
	for _, f := range modules {
		f(r, ws)
	}

	log.Infof("loaded %d modules", len(modules))
	r.Use(static.Serve("/", static.LocalFile(ws.Webroot, false)))
}
