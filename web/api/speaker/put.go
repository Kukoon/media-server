package speaker

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

// @Summary Edit Speaker
// @Description Edit speaker by ID
// @Tags speaker
// @Produce  json
// @Success 200 {object} models.Speaker
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/speaker/{speaker_id} [put]
// @Param speaker_id path string false "uuid of speaker"
// @Param body body models.Speaker false "new values in speaker"
// @Security ApiKeyAuth
func apiPut(r *gin.Engine, ws *web.Service) {
	r.PUT("/api/v1/speaker/:uuid", auth.MiddlewarePermissionParamUUID(ws, models.Speaker{}), func(c *gin.Context) {
		old := models.Speaker{
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
		var data models.Speaker
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, web.HTTPError{
				Message: web.ErrAPIInvalidRequestFormat.Error(),
				Error:   err.Error(),
			})
			return
		}

		data.ID = old.ID
		data.OwnerID = old.OwnerID

		if err := ws.DB.Omit("Owner.*").Save(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &data)
	})
}
