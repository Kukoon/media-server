package recording

import (
	"fmt"
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kukoon/media-server/models"
)

// @Summary List all Recordings
// @Description Show a list of all recordings
// @Produce  json
// @Success 200 {array} models.Recording
// @Failure 400 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/recordings [get]
// @Param channel query string false "filter by UUID of a channel"
// @Param event query string false "filter by UUID of a event"
// @Param tag query string false "filter by UUID of any tag (multiple times)"
// @Param speaker query string false "filter by UUID of any speaker (multiple times)"
// @Param lang query string false "show description in given language"
func apiList(r *gin.Engine, ws *web.Service) {
	r.GET("/api/v1/recordings", func(c *gin.Context) {
		list := []*models.Recording{}
		db := ws.DB

		// channel
		db = db.Joins("Channel")
		if str, ok := c.GetQuery("channel"); ok {
			uuid, err := uuid.Parse(str)
			if err != nil {
				c.JSON(http.StatusBadRequest, web.HTTPError{
					Message: web.ErrAPIInvalidRequestFormat.Error(),
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
					Message: web.ErrAPIInvalidRequestFormat.Error(),
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
			for i, str := range strArray {
				id, err := uuid.Parse(str)
				if err != nil {
					c.JSON(http.StatusBadRequest, web.HTTPError{
						Message: web.ErrAPIInvalidRequestFormat.Error(),
						Error:   err.Error(),
					})
					return
				}
				db = db.Joins(fmt.Sprintf("INNER JOIN recording_tags rt%d ON rt%d.recording_id = recordings.id AND rt%d.tag_id = ?", i, i, i), id)
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
				db = db.Joins(fmt.Sprintf("INNER JOIN recording_speakers rs%d ON rs%d.recording_id = recordings.id AND rs%d.speaker_id = ?", i, i, i), id)
			}
		}

		// TODO login - own channel
		db = db.Where("public", true).Where("listed", true)

		if err := db.Order("created_at DESC").Find(&list).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &list)
	})
}
