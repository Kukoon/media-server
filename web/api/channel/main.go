package channel

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kukoon/media-server/models"
	"github.com/Kukoon/media-server/web"
)

func init() {
	web.ModuleRegister(func(r *gin.Engine, ws *web.Service) {
		// @Summary List all Channels
		// @Description Show a list of all channels
		// @Produce  json
		// @Success 200 {array} models.Channel
		// @Failure 500 {object} HTTPError
		// @Router /api/v1/channels [get]
		r.GET("/api/v1/channels", func(c *gin.Context) {
			list := []*models.Channel{}
			if err := ws.DB.Find(&list).Error; err != nil {
				c.JSON(http.StatusInternalServerError, web.HTTPError{
					Message: web.APIErrorInternalDatabase,
					Error:   err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, &list)
		})

		// @Summary Show Channel
		// @Description Show channel with all informations
		// @Produce  json
		// @Success 200 {object} models.Channel
		// @Failure 400 {object} HTTPError
		// @Failure 404 {object} HTTPError
		// @Router /api/v1/channel/{slug} [get]
		// @Param slug path string false "slug or uuid of channel"
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
						Message: web.APIErrorNotFound,
						Error:   err.Error(),
					})
					c.JSON(http.StatusNotFound, err.Error())
					return
				}
				c.JSON(http.StatusInternalServerError, web.HTTPError{
					Message: web.APIErrorInternalDatabase,
					Error:   err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, &obj)
		})
	})
}
