package channel

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/Kukoon/media-server/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Add Channel
// @Description Add channel with owner you
// @Tags channel
// @Produce  json
// @Success 200 {object} models.Channel
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/channel [post]
// @Param body body models.Channel false "channel"
// @Security ApiKeyAuth
func apiPost(r *gin.Engine, ws *web.Service) {
	r.POST("/api/v1/channel", func(c *gin.Context) {
		id, ok := auth.GetCurrentUserID(c)
		if !ok {
			return
		}

		var data models.Channel
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, web.HTTPError{
				Message: web.ErrAPIInvalidRequestFormat.Error(),
				Error:   err.Error(),
			})
			return
		}

		data.ID = uuid.Nil
		data.Owners = []models.User{{ID: id}}

		if err := ws.DB.Omit("Owners.*", "Recordings").Create(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	})
}
