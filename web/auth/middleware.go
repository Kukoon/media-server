package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/Kukoon/media-server/web"
)

func MiddlewareLogin(ws *web.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := GetCurrentUserID(c)
		if !ok {
			c.Abort()
		}
	}
}
