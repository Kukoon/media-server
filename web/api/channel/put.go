package channel

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Kukoon/media-server/models"
)

// @Summary Save Channel
// @Description Save new data for given Channel
// @Tags channel
// @Produce  json
// @Success 200 {object} models.Channel
// @Failure 400 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Router /api/v1/channel/{channel_id} [put]
// @Param channel_id path string false "uuid of channel"
// @Param body body models.Channel false "new values in channel"
// @Security ApiKeyAuth
func apiPut(r *gin.Engine, ws *web.Service) {
	r.PUT("/api/v1/channel/:slug", auth.MiddlewarePermissionParam(ws, models.Channel{}, "slug"), func(c *gin.Context) {
		id := uuid.MustParse(c.Params.ByName("slug"))
		var data models.Channel
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, web.HTTPError{
				Message: web.ErrAPIInvalidRequestFormat.Error(),
				Error:   err.Error(),
			})
			return
		}
		data.ID = id
		data.Recordings = []*models.Recording{}
		if err := ws.DB.Omit("Owner.*", "Recordings").Save(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &data)
	})
}
