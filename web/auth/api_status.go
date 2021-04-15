package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Kukoon/media-server/web"
)

// @Summary Login status
// @Description show user_id and username if logged in
// @Accept json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 401 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/auth/status [get]
func init() {
	web.ModuleRegister(func(r *gin.Engine, ws *web.Service) {
		r.GET("/api/v1/auth/status", MiddlewareLogin(ws), func(c *gin.Context) {
			d, ok := GetCurrentUser(c, ws)
			if ok {
				c.JSON(http.StatusOK, d)
			}
		})
	})
}
