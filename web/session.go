package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func (config *Service) LoadSession(r *gin.Engine) {
	store := cookie.NewStore([]byte(config.Session.Secret))
	r.Use(sessions.Sessions(config.Session.Name, store))
}
