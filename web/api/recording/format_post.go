package recording

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Kukoon/media-server/models"
)

// @Summary Create new recording format
// @Description Create new recording format on given Recording
// @Tags recording
// @Produce  json
// @Success 200 {object} models.RecordingFormat
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/recording/{recording_id}/format [post]
// @Param recording_id path string false "uuid of recording"
// @Param body body models.RecordingFormat false "new values in recording description"
// @Security ApiKeyAuth
func apiFormatPost(r *gin.Engine, ws *web.Service) {
	r.POST("/api/v1/recording/:uuid/format", auth.MiddlewarePermissionParamUUID(ws, models.Recording{}), func(c *gin.Context) {
		id := uuid.MustParse(c.Params.ByName("uuid"))
		var data models.RecordingFormat
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, web.HTTPError{
				Message: web.ErrAPIInvalidRequestFormat.Error(),
				Error:   err.Error(),
			})
			return
		}

		data.ID = uuid.Nil
		data.RecordingID = id

		if err := ws.DB.Omit("CreatedAt", "UpdatedAt").Create(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &data)
	})
}
