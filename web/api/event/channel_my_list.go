package event

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Kukoon/media-server/models"
)

// @Summary List Events of my Channel
// @Description Show a list of all events on a given channel (with my permission)
// @Tags event
// @Produce  json
// @Success 200 {array} models.Event
// @Failure 400 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/channel/{slug}/events [get]
// @Param slug path string false "uuid of channel"
// @Security ApiKeyAuth
func apiChannelListMy(r *gin.Engine, ws *web.Service) {
	r.GET("/api/v1/channel/:slug/events", auth.MiddlewarePermissionParam(ws, models.Channel{}, "slug"), func(c *gin.Context) {
		list := []*models.Event{}
		if err := ws.DB.
			Where("owner_id = ?", uuid.MustParse(c.Params.ByName("slug"))).
			Find(&list).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &list)
	})
}
