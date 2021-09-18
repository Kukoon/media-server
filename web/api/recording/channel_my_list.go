package recording

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Kukoon/media-server/models"
)

// @Summary List Recordings of my Channel
// @Description Show a list of all recordings on a given channel (with my permission)
// @Tags recording
// @Produce  json
// @Success 200 {array} models.Recording
// @Failure 400 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/channel/{slug}/recordings [get]
// @Param slug path string false "uuid of channel"
// @Param event query string false "filter by UUID of a event"
// @Param tag query string false "filter by UUID of any tag (multiple times)"
// @Param speaker query string false "filter by UUID of any speaker (multiple times)"
// @Param lang query string false "show description in given language"
func apiChannelListMy(r *gin.Engine, ws *web.Service) {
	r.GET("/api/v1/channel/:slug/recordings", auth.MiddlewarePermissionParam(ws, models.Channel{}, "slug"), func(c *gin.Context) {
		db, ok := filterRecordings(ws.DB, c)
		if !ok {
			return
		}

		list := []*models.Recording{}
		// TODO no filter for listen_at
		if err := db.
			Where("channel_id = ?", uuid.MustParse(c.Params.ByName("slug"))).
			Order("created_at DESC").
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
