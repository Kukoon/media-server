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

// @Summary Edit Recording Description
// @Description Edit recording description by ID
// @Tags recording
// @Produce  json
// @Success 200 {object} models.RecordingLang
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/recording-lang/{lang_id} [put]
// @Param lang_id path string false "uuid of recording"
// @Param body body models.RecordingLang false "new values in recording"
// @Security ApiKeyAuth
func apiLangPut(r *gin.Engine, ws *web.Service) {
	r.PUT("/api/v1/recording-lang/:uuid", auth.MiddlewarePermissionParamUUID(ws, models.RecordingLang{}), func(c *gin.Context) {
		old := models.RecordingLang{
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
		var data models.RecordingLang
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
