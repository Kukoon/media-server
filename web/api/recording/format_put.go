package recording

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

// @Summary Edit Recording Format
// @Description Edit recording format by ID
// @Tags recording
// @Produce  json
// @Success 200 {object} models.RecordingFormat
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/recording-format/{format_id} [put]
// @Param format_id path string false "uuid of recording"
// @Param body body models.RecordingFormat false "new values in recording"
// @Security ApiKeyAuth
func apiFormatPut(r *gin.Engine, ws *web.Service) {
	r.PUT("/api/v1/recording-format/:uuid", auth.MiddlewarePermissionParamUUID(ws, models.RecordingFormat{}), func(c *gin.Context) {
		old := models.RecordingFormat{
			ID: uuid.MustParse(c.Params.ByName("uuid")),
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
		var data models.RecordingFormat
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, web.HTTPError{
				Message: web.ErrAPIInvalidRequestFormat.Error(),
				Error:   err.Error(),
			})
			return
		}

		data.ID = old.ID
		data.RecordingID = old.RecordingID

		if err := ws.DB.Omit("CreatedAt", "UpdatedAt").Save(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &data)
	})
}
