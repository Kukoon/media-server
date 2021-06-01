package stream

import (
	"net/http"
	"strconv"

	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kukoon/media-server/models"
)

// @Summary List all Streams
// @Description Show a list of all streams
// @Produce  json
// @Success 200 {array} models.PublicStream
// @Failure 400 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/streams [get]
// @Param running query bool false "filter by running streams"
// @Param channel query string false "filter by UUID of a channel"
// @Param event query string false "filter by UUID of a event"
// @Param tag query string false "filter by UUID of any tag (multiple times)"
// @Param speaker query string false "filter by UUID of any speaker (multiple times)"
// @Param lang query string false "show description in given language"
func init() {
	web.ModuleRegister(func(r *gin.Engine, ws *web.Service) {
		r.GET("/api/v1/streams", func(c *gin.Context) {
			list := []*models.Stream{}
			db := ws.DB

			// running
			if str, ok := c.GetQuery("running"); ok {
				b, err := strconv.ParseBool(str)
				if err != nil {
					c.JSON(http.StatusBadRequest, web.HTTPError{
						Message: web.APIErrorInvalidRequestFormat,
						Error:   err.Error(),
					})
					return
				}
				db = db.Where("running", b)
			}

			// channel
			db = db.Joins("Channel")
			if str, ok := c.GetQuery("channel"); ok {
				uuid, err := uuid.Parse(str)
				if err != nil {
					c.JSON(http.StatusBadRequest, web.HTTPError{
						Message: web.APIErrorInvalidRequestFormat,
						Error:   err.Error(),
					})
					return
				}
				db = db.Where("channel_id", uuid)
			}

			// event
			db = db.Joins("Event")
			if str, ok := c.GetQuery("event"); ok {
				uuid, err := uuid.Parse(str)
				if err != nil {
					c.JSON(http.StatusBadRequest, web.HTTPError{
						Message: web.APIErrorInvalidRequestFormat,
						Error:   err.Error(),
					})
					return
				}
				db = db.Where("event_id", uuid)
			}

			// tags + language
			if str, ok := c.GetQuery("lang"); ok {
				db = db.Preload("Lang", func(db *gorm.DB) *gorm.DB {
					return db.Where("lang", str)
				}).Preload("Tags.Lang", func(db *gorm.DB) *gorm.DB {
					return db.Where("lang", str)
				})
			} else {
				db = db.Preload("Tags")
			}
			// filter tag
			if strArray, ok := c.GetQueryArray("tag"); ok {
				ids := []uuid.UUID{}
				for _, str := range strArray {
					id, err := uuid.Parse(str)
					if err != nil {
						c.JSON(http.StatusBadRequest, web.HTTPError{
							Message: web.APIErrorInvalidRequestFormat,
							Error:   err.Error(),
						})
						return
					}
					ids = append(ids, id)
				}
				db = db.Joins("LEFT JOIN stream_tags ON stream_tags.stream_id = streams.id").Where("tag_id IN (?)", ids)
			}
			// filter speaker
			db = db.Preload("Speakers")
			if strArray, ok := c.GetQueryArray("speaker"); ok {
				ids := []uuid.UUID{}
				for _, str := range strArray {
					id, err := uuid.Parse(str)
					if err != nil {
						c.JSON(http.StatusBadRequest, web.HTTPError{
							Message: web.APIErrorInvalidRequestFormat,
							Error:   err.Error(),
						})
						return
					}
					ids = append(ids, id)
				}
				db = db.Joins("LEFT JOIN stream_speakers ON stream_speakers.stream_id = streams.id").Where("speaker_id IN (?)", ids)
			}
			if err := db.Find(&list).Error; err != nil {
				c.JSON(http.StatusInternalServerError, web.HTTPError{
					Message: web.APIErrorInternalDatabase,
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
	})
}
