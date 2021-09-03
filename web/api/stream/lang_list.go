package stream

import (
	"errors"
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kukoon/media-server/models"
)

// @Summary List all Language Description of an Stream
// @Description List all Descriptions/Languages of stream by ID
// @Tags stream
// @Produce  json
// @Success 200 {array} models.StreamLang
// @Failure 400 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/stream/{stream_id}/langs [get]
// @Param stream_id path string false "uuid of stream"
// @Param lang query string false "show description in given language"
func apiLangList(r *gin.Engine, ws *web.Service) {
	r.GET("/api/v1/stream/:uuid/langs", func(c *gin.Context) {
		id, err := uuid.Parse(c.Params.ByName("uuid"))
		if err != nil {
			c.JSON(http.StatusBadRequest, web.HTTPError{
				Message: web.ErrAPIInvalidRequestFormat.Error(),
				Error:   err.Error(),
			})
			return
		}
		db := ws.DB.Where("stream_id = ?", id)
		if str, ok := c.GetQuery("lang"); ok {
			db = db.Where("lang", str)
		}

		list := []*models.StreamLang{}
		if err := db.Find(&list).Error; err != nil {
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

		c.JSON(http.StatusOK, &list)
	})
}
