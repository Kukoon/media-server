package stream

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Kukoon/media-server/models"
)

// @Summary Create new stream description
// @Description Create new stream description on given Stream
// @Tags stream
// @Produce  json
// @Success 200 {object} models.StreamLang
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/stream/{stream_id}/lang [post]
// @Param stream_id path string false "uuid of stream"
// @Param body body models.StreamLang false "new values in stream description"
// @Security ApiKeyAuth
func apiLangPost(r *gin.Engine, ws *web.Service) {
	r.POST("/api/v1/stream/:uuid/lang", auth.MiddlewarePermissionParamUUID(ws, models.Stream{}), func(c *gin.Context) {
		id := uuid.MustParse(c.Params.ByName("uuid"))
		var data models.StreamLang
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, web.HTTPError{
				Message: web.ErrAPIInvalidRequestFormat.Error(),
				Error:   err.Error(),
			})
			return
		}

		data.ID = uuid.Nil
		data.StreamID = id

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
