package channel

import (
	"errors"
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kukoon/media-server/models"
)

// @Summary Show Channel
// @Description Show channel with all informations
// @Tags channel
// @Produce  json
// @Success 200 {object} models.Channel
// @Failure 400 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Router /api/v1/channel/{slug} [get]
// @Param slug path string false "slug or uuid of channel"
func apiGet(r *gin.Engine, ws *web.Service) {
	r.GET("/api/v1/channel/:slug", func(c *gin.Context) {
		slug := c.Params.ByName("slug")
		db := ws.DB
		obj := models.Channel{}

		uuid, err := uuid.Parse(slug)
		if err != nil {
			db = db.Where("common_name", slug)
			obj.CommonName = slug
		} else {
			obj.ID = uuid
		}
		if err := db.First(&obj).Error; err != nil {
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

		c.JSON(http.StatusOK, &obj)
	})
}
