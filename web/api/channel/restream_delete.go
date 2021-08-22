package channel

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	ovenAPI "dev.sum7.eu/genofire/oven-exporter/api"
	"github.com/Kukoon/media-server/models"
	"github.com/Kukoon/media-server/oven"
	"github.com/gin-gonic/gin"
)

// @Summary Delete Restream of Channel
// @Description Delete all restream / push of channel
// @Tags channel
// @Produce  json
// @Success 200 {object} Restream
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/channel/{channel_id}/restream/{id} [delete]
func apiRestreamDelete(r *gin.Engine, ws *web.Service, oven *oven.Service) {
	r.DELETE("/api/v1/channel/:slug/restream/:id", auth.MiddlewarePermissionParam(ws, models.Channel{}, "slug"), func(c *gin.Context) {
		channelid := c.Params.ByName("slug")
		id := c.Params.ByName("id")
		resp, err := oven.Client.RequestPushStatusDefault()
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}
		var obj *ovenAPI.ResponsePushData
		for _, data := range resp.Data {
			if data.ID == id {
				obj = data
				break
			}
		}
		if obj == nil {
			c.JSON(http.StatusNotFound, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
			})
			return
		} else if obj.Stream.Name != channelid {
			c.JSON(http.StatusBadRequest, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
			})
			return
		}

		if err := oven.Client.DeletePushDefault(id); err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, RestreamFromOven(obj))
	})
}
