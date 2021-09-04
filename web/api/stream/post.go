package stream

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Kukoon/media-server/models"
)

// @Summary Create new stream metadata
// @Description Create new stream metadata on given Channel
// @Tags stream
// @Produce  json
// @Success 200 {object} models.Stream
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/channel/{channel_id}/stream [post]
// @Param channel_id path string false "uuid of channel"
// @Param body body Stream false "new values in stream"
// @Security ApiKeyAuth
func apiPost(r *gin.Engine, ws *web.Service) {
	r.POST("/api/v1/channel/:slug/stream", auth.MiddlewarePermissionParam(ws, models.Channel{}, "slug"), func(c *gin.Context) {
		id := uuid.MustParse(c.Params.ByName("slug"))
		var req Stream
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.HTTPError{
				Message: web.ErrAPIInvalidRequestFormat.Error(),
				Error:   err.Error(),
			})
			return
		}

		data := req.Model()
		data.ID = uuid.Nil
		data.ChannelID = id

		if err := ws.DB.Omit("Lang", "Tags.*", "Speakers.*").Create(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &data)
	})
}
