package event

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/gin-gonic/gin"

	"github.com/Kukoon/media-server/models"
)

// @Summary List all Events
// @Description Show a list of all events
// @Produce  json
// @Success 200 {array} models.Event
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/events [get]
func apiList(r *gin.Engine, ws *web.Service) {
	r.GET("/api/v1/events", func(c *gin.Context) {
		list := []*models.Event{}
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
