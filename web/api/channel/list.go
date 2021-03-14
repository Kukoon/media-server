package channel

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Kukoon/media-server/models"
	"github.com/Kukoon/media-server/web"
)

// @Summary List all Channels
// @Description Show a list of all channels
// @Produce  json
// @Success 200 {array} models.Channel
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/channels [get]
func init() {
	web.ModuleRegister(func(r *gin.Engine, ws *web.Service) {
		r.GET("/api/v1/channels", func(c *gin.Context) {
			list := []*models.Channel{}
			if err := ws.DB.Find(&list).Error; err != nil {
				c.JSON(http.StatusInternalServerError, web.HTTPError{
					Message: web.APIErrorInternalDatabase,
					Error:   err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, &list)
		})
	})
}
