package speaker

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Kukoon/media-server/models"
)

// @Summary Create new speaker
// @Description Create new speaker on given Channel
// @Tags speaker
// @Produce  json
// @Success 200 {object} models.Speaker
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/channel/{channel_id}/speaker [post]
// @Param channel_id path string false "uuid of channel"
// @Param body body models.Speaker false "values for speaker"
// @Security ApiKeyAuth
func apiPost(r *gin.Engine, ws *web.Service) {
	r.POST("/api/v1/channel/:slug/speaker", auth.MiddlewarePermissionParam(ws, models.Channel{}, "slug"), func(c *gin.Context) {
		id := uuid.MustParse(c.Params.ByName("slug"))
		var data models.Speaker
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
