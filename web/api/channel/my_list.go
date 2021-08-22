package channel

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"

	"github.com/Kukoon/media-server/models"
)

// @Summary List my Channels
// @Description Show a list of all channels you has permission
// @Produce  json
// @Success 200 {array} models.Channel
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/my/channels [get]
func apiListMy(r *gin.Engine, ws *web.Service) {
	r.GET("/api/v1/my/channels", func(c *gin.Context) {
		id, ok := auth.GetCurrentUserID(c)
		if !ok {
			return
		}
		list := []*models.Channel{}
		if err := ws.DB.
			Joins("INNER JOIN user_channels uc ON uc.channel_id=channels.id AND uc.user_id= ?", id).
			Find(&list).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &list)
	})
}
