package stream

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func filterStreams(tx *gorm.DB, c *gin.Context) (db *gorm.DB, isOK bool) {
	now := time.Now()
	db = tx

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
	isOK = true
	return
}
