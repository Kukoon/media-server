package stream

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Kukoon/media-server/models"
)

// @Summary List Streams of my Channel
// @Description Show a list of all streams on a given channel (with my permission)
// @Tags stream
// @Produce  json
// @Success 200 {array} models.Stream
// @Failure 400 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/channel/{slug}/streams [get]
// @Param slug path string false "slug or uuid of stream"
// @Param running query bool false "filter by running streams"
// @Param upcoming query bool false "filter by next streams"
// @Param from query bool false "filter by date start streams"
// @Param to query bool false "filter by date end streams"
// @Param event query string false "filter by UUID of a event"
// @Param tag query string false "filter by UUID of any tag (multiple times)"
// @Param speaker query string false "filter by UUID of any speaker (multiple times)"
func apiChannelListMy(r *gin.Engine, ws *web.Service) {
	r.GET("/api/v1/channel/:slug/streams", auth.MiddlewarePermissionParam(ws, models.Channel{}, "slug"), func(c *gin.Context) {
		list := []*models.Stream{}
		db := ws.DB
		now := time.Now()

		// running
		if str, ok := c.GetQuery("running"); ok {
			b, err := strconv.ParseBool(str)
			if err != nil {
				c.JSON(http.StatusBadRequest, web.HTTPError{
					Message: web.ErrAPIInvalidRequestFormat.Error(),
					Error:   err.Error(),
				})
				return
			}
			db = db.Where("running", b)
		}
		// upcoming
		if str, ok := c.GetQuery("upcoming"); ok {
			b, err := strconv.ParseBool(str)
			if err != nil {
				c.JSON(http.StatusBadRequest, web.HTTPError{
					Message: web.ErrAPIInvalidRequestFormat.Error(),
					Error:   err.Error(),
				})
				return
			}
			if b {
				db = db.Where("start_at > ?", now)
			} else {
				db = db.Where("start_at <= ?", now)
			}
			// TODO - here order?
			db = db.Order("start_at")
		}
		// from
		if str, ok := c.GetQuery("from"); ok {
			t, err := time.Parse(time.RFC3339, str)
			if err != nil {
				c.JSON(http.StatusBadRequest, web.HTTPError{
					Message: web.ErrAPIInvalidRequestFormat.Error(),
					Error:   err.Error(),
				})
				return
			}
			db = db.Where("start_at >= ?", t)
		}
		// to
		if str, ok := c.GetQuery("to"); ok {
			t, err := time.Parse(time.RFC3339, str)
			if err != nil {
				c.JSON(http.StatusBadRequest, web.HTTPError{
					Message: web.ErrAPIInvalidRequestFormat.Error(),
					Error:   err.Error(),
				})
				return
			}
			db = db.Where("start_at <= ?", t)
		}

		// event
		db = db.Joins("Event")
		if str, ok := c.GetQuery("event"); ok {
			uuid, err := uuid.Parse(str)
			if err != nil {
				c.JSON(http.StatusBadRequest, web.HTTPError{
					Message: web.ErrAPIInvalidRequestFormat.Error(),
					Error:   err.Error(),
				})
				return
			}
			db = db.Where("event_id", uuid)
		}

		// filter tag
		db = db.Preload("Tags")
		if strArray, ok := c.GetQueryArray("tag"); ok {
			for i, str := range strArray {
				id, err := uuid.Parse(str)
				if err != nil {
					c.JSON(http.StatusBadRequest, web.HTTPError{
						Message: web.ErrAPIInvalidRequestFormat.Error(),
						Error:   err.Error(),
					})
					return
				}
				db = db.Joins(fmt.Sprintf("INNER JOIN stream_tags st%d ON st%d.stream_id = streams.id AND st%d.tag_id = ?", i, i, i), id)
			}
		}
		// filter speaker
		db = db.Preload("Speakers")
		if strArray, ok := c.GetQueryArray("speaker"); ok {
			for i, str := range strArray {
				id, err := uuid.Parse(str)
				if err != nil {
					c.JSON(http.StatusBadRequest, web.HTTPError{
						Message: web.ErrAPIInvalidRequestFormat.Error(),
						Error:   err.Error(),
					})
					return
				}
				db = db.Joins(fmt.Sprintf("INNER JOIN stream_speakers ss%d ON ss%d.stream_id = streams.id AND ss%d.speaker_id = ?", i, i, i), id)
			}
		}
		// TODO no filter for listen_at
		if err := db.
			Joins("Channel").
			Where("channel_id = ?", uuid.MustParse(c.Params.ByName("slug"))).
			Order("start_at").
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
