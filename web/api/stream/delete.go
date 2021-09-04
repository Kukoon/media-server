package stream

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Kukoon/media-server/models"
)

// @Summary Delete Stream
// @Description Delete Stream
// @Tags stream
// @Produce  json
// @Success 200 {object} bool "true if deleted"
// @Failure 400 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Router /api/v1/stream/{stream_id} [delete]
// @Param stream_id path string false "uuid of stream"
// @Security ApiKeyAuth
func apiDelete(r *gin.Engine, ws *web.Service) {
	r.DELETE("/api/v1/stream/:uuid", auth.MiddlewarePermissionParamUUID(ws, models.Stream{}), func(c *gin.Context) {
		id := uuid.MustParse(c.Params.ByName("uuid"))
		if err := ws.DB.Delete(&models.Stream{ID: id}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, true)
	})
}
