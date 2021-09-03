package stream

import (
	"net/http"
	"time"

	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/gin-gonic/gin"

	"github.com/Kukoon/media-server/models"
)

// @Summary List all Streams
// @Description Show a list of all streams
// @Tags stream
// @Produce  json
// @Success 200 {array} models.PublicStream
// @Failure 400 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/streams [get]
// @Param running query bool false "filter by running streams"
// @Param upcoming query bool false "filter by next streams"
// @Param from query bool false "filter by date start streams"
// @Param to query bool false "filter by date end streams"
// @Param channel query string false "filter by UUID of a channel (multiple times)"
// @Param event query string false "filter by UUID of a event"
// @Param tag query string false "filter by UUID of any tag (multiple times)"
// @Param speaker query string false "filter by UUID of any speaker (multiple times)"
// @Param lang query string false "show description in given language"
func apiList(r *gin.Engine, ws *web.Service) {
	r.GET("/api/v1/streams", func(c *gin.Context) {
		db, ok := filterStreams(ws.DB, c)
		if !ok {
			return
		}

		list := []*models.Stream{}
		now := time.Now()

		/* channel
		db = db.Joins("Channel")
		if strArray, ok := c.GetQueryArray("channel"); ok {
			ids := make([]uuid.UUID, len(strArray))
			for i, str := range strArray {
				id, err := uuid.Parse(str)
				if err != nil {
					c.JSON(http.StatusBadRequest, web.HTTPError{
						Message: web.ErrAPIInvalidRequestFormat.Error(),
						Error:   err.Error(),
					})
					return
				}
				ids[i] = id
			}
			db = db.Where("channel_id IN (?)", ids)
		}
		*/

		// TODO - here order?
		if err := db.
			Where("listen_at < ?", now).
			Order("start_at").
			Find(&list).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		listOutput := []*models.PublicStream{}
		for _, obj := range list {
			listOutput = append(listOutput, obj.GetPublic())
		}

		c.JSON(http.StatusOK, &listOutput)
	})
}
