package channel

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/Kukoon/media-server/oven"
	"github.com/gin-gonic/gin"
)

// @Summary List Restreams of Channel
// @Description Show a list of all restream / push of channel
// @Produce  json
// @Success 200 {array} models.Channel
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/channel/{id}/restreams [get]
func apiRestreamList(r *gin.Engine, oven *oven.Service) {
	r.GET("/api/v1/channel/:slug/restreams", func(c *gin.Context) {
		id := c.Params.ByName("slug")
		resp, err := oven.Client.RequestPushStatusDefault()
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}
		list := []*Restream{}
		for _, data := range resp.Data {
			if data.Stream.Name == id {
				list = append(list, RestreamFromOven(data))
			}
		}
		c.JSON(http.StatusOK, &list)
	})
}
