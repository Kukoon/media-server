package stream

import (
	"errors"
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kukoon/media-server/models"
)

// @Summary Edit Stream Metadata
// @Description Edit stream by ID
// @Tags stream
// @Produce  json
// @Success 200 {object} models.Stream
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/stream/{stream_id} [put]
// @Param stream_id path string false "uuid of stream"
// @Param body body Stream false "new values in stream"
// @Security ApiKeyAuth
func apiPut(r *gin.Engine, ws *web.Service) {
	r.PUT("/api/v1/stream/:sid", auth.MiddlewarePermissionParam(ws, models.Stream{}, "sid"), func(c *gin.Context) {
		old := models.Stream{
			ID: uuid.MustParse(c.Params.ByName("sid")),
		}
		if err := ws.DB.First(&old).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, web.HTTPError{
					Message: web.ErrAPINotFound.Error(),
					Error:   err.Error(),
				})
				return
			}
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}
		var req Stream
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.HTTPError{
				Message: web.ErrAPIInvalidRequestFormat.Error(),
				Error:   err.Error(),
			})
			return
		}

		data := req.Model()
		data.ID = old.ID
		data.ChannelID = old.ChannelID

		if err := ws.DB.Omit("Lang", "Tags.*", "Speakers.*").Save(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &data)
	})
}
