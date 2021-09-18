package event

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Kukoon/media-server/models"
)

// @Summary Create new event
// @Description Create new event on given Channel
// @Tags event
// @Produce  json
// @Success 200 {object} models.Event
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/channel/{channel_id}/event [post]
// @Param channel_id path string false "uuid of channel"
// @Param body body models.Event false "values for event"
// @Security ApiKeyAuth
func apiPost(r *gin.Engine, ws *web.Service) {
	r.POST("/api/v1/channel/:slug/event", auth.MiddlewarePermissionParam(ws, models.Channel{}, "slug"), func(c *gin.Context) {
		id := uuid.MustParse(c.Params.ByName("slug"))
		var data models.Event
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, web.HTTPError{
				Message: web.ErrAPIInvalidRequestFormat.Error(),
				Error:   err.Error(),
			})
			return
		}

		data.ID = uuid.Nil
		data.OwnerID = id

		if err := ws.DB.Omit("Owner.*").Create(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &data)
	})
}
