package recording

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Kukoon/media-server/models"
)

// @Summary Delete Recording
// @Description Delete Recording
// @Tags recording
// @Produce  json
// @Success 200 {object} bool "true if deleted"
// @Failure 400 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/recording/{recording_id} [delete]
// @Param recording_id path string false "uuid of recording"
// @Security ApiKeyAuth
func apiDelete(r *gin.Engine, ws *web.Service) {
	r.DELETE("/api/v1/recording/:uuid", auth.MiddlewarePermissionParamUUID(ws, models.Recording{}), func(c *gin.Context) {
		id := uuid.MustParse(c.Params.ByName("uuid"))
		result := ws.DB.Delete(&models.Recording{ID: id})
		if err := result.Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		if result.RowsAffected < 1 {
			c.JSON(http.StatusNotFound, web.HTTPError{
				Message: web.ErrAPINotFound.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, result.RowsAffected == 1)
	})
}
