package channel

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/gin-gonic/gin"

	"github.com/Kukoon/media-server/models"
)

// @Summary List all Channels
// @Description Show a list of all channels
// @Tags channel
// @Produce  json
// @Success 200 {array} models.Channel
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/channels [get]
func apiList(r *gin.Engine, ws *web.Service) {
	r.GET("/api/v1/channels", func(c *gin.Context) {
		list := []*models.Channel{}
		if err := ws.DB.Find(&list).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &list)
	})
}
