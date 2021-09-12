package recording

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Kukoon/media-server/models"
)

// @Summary Create new recording metadata
// @Description Create new recording metadata on given Channel
// @Tags recording
// @Produce  json
// @Success 200 {object} models.Recording
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/channel/{channel_id}/recording [post]
// @Param channel_id path string false "uuid of channel"
// @Param body body Recording false "new values in recording"
// @Security ApiKeyAuth
func apiPost(r *gin.Engine, ws *web.Service) {
	r.POST("/api/v1/channel/:slug/recording", auth.MiddlewarePermissionParam(ws, models.Channel{}, "slug"), func(c *gin.Context) {
		id := uuid.MustParse(c.Params.ByName("slug"))
		var req Recording
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

		if err := ws.DB.Omit("Format", "Lang", "Tags.*", "Speakers.*").Create(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &data)
	})
}
