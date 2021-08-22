package channel

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/Kukoon/media-server/models"
	"github.com/Kukoon/media-server/oven"
	"github.com/bdlm/log"
	"github.com/gin-gonic/gin"
)

// @Summary Add Restream to Channel
// @Description Add restream / push to channel
// @Produce  json
// @Success 200 {object} Restream
// @Failure 400 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/channel/{channel_id}/restream [post]
func apiRestreamAdd(r *gin.Engine, ws *web.Service, oven *oven.Service) {
	r.POST("/api/v1/channel/:slug/restream", auth.MiddlewarePermissionParam(ws, models.Channel{}, "slug"), func(c *gin.Context) {
		channelid := c.Params.ByName("slug")

		var data RestreamAdd
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, web.HTTPError{
				Message: web.ErrAPIInvalidRequestFormat.Error(),
				Error:   err.Error(),
			})
			return
		}

		resp, err := oven.Client.StartPushDefault(data.ToOven(channelid))
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		log.Error(len(resp.Data))

		c.JSON(http.StatusOK, RestreamFromOven(resp.Data[0]))
	})
}
