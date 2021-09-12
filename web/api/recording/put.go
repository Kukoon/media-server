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

// @Summary Edit Recording Metadata
// @Description Edit recording by ID
// @Tags recording
// @Produce  json
// @Success 200 {object} models.Recording
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/recording/{recording_id} [put]
// @Param recording_id path string false "uuid of recording"
// @Param body body Recording false "new values in recording"
// @Security ApiKeyAuth
func apiPut(r *gin.Engine, ws *web.Service) {
	r.PUT("/api/v1/recording/:uuid", auth.MiddlewarePermissionParamUUID(ws, models.Recording{}), func(c *gin.Context) {
		old := models.Recording{
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
		var req Recording
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

		if err := ws.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Omit("Format", "Lang", "Tags.*", "Speakers.*").Save(&data).Error; err != nil {
				return err
			}
			if err := tx.Model(&data).Association("Tags").Replace(data.Tags); err != nil {
				return err
			}
			if err := tx.Model(&data).Association("Speakers").Replace(data.Speakers); err != nil {
				return err
			}
			return nil
		}); err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		ws.DB.Preload("Event").First(&data)

		c.JSON(http.StatusOK, &data)
	})
}
