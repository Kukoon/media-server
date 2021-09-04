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

// @Summary Get Stream Metadata
// @Description Get stream Metadata by ID
// @Tags stream
// @Produce  json
// @Success 200 {object} models.Stream
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/stream/{stream_id} [get]
// @Param stream_id path string false "uuid of stream"
// @Security ApiKeyAuth
func apiGet(r *gin.Engine, ws *web.Service) {
	r.GET("/api/v1/stream/:uuid", auth.MiddlewarePermissionParamUUID(ws, models.Stream{}), func(c *gin.Context) {
		data := models.Stream{
			ID: uuid.MustParse(c.Params.ByName("uuid")),
		}

		if err := ws.DB.First(&data).Error; err != nil {
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

		c.JSON(http.StatusOK, &data)
	})
}
