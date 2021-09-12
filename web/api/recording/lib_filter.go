package recording

import (
	"fmt"
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func filterRecordings(tx *gorm.DB, c *gin.Context) (db *gorm.DB, isOK bool) {
	db = tx

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

	// channel
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

	isOK = true
	return
}
