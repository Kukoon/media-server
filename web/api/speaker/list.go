package speaker

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/gin-gonic/gin"

	"github.com/Kukoon/media-server/models"
)

// @Summary List all Speakers
// @Description Show a list of all speakers
// @Produce  json
// @Success 200 {array} models.Speaker
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/speakers [get]
func apiList(r *gin.Engine, ws *web.Service) {
	r.GET("/api/v1/speakers", func(c *gin.Context) {
		list := []*models.Speaker{}
		if err := ws.DB.Order("name").Find(&list).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &list)
	})
}
