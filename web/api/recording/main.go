package recording

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kukoon/media-server/models"
	"github.com/Kukoon/media-server/web"
)

func init() {
	web.ModuleRegister(func(r *gin.Engine, ws *web.Service) {

		// @Summary List all Recordings
		// @Description Show a list of all recordings
		// @Produce  json
		// @Success 200 {array} models.Recording
		// @Failure 400 {object} HTTPError
		// @Failure 500 {object} HTTPError
		// @Router /api/v1/recordings [get]
		// @Param channel query string false "filter by UUID of a channel"
		// @Param event query string false "filter by UUID of a event"
		// @Param tag query string false "filter by UUID of any tag (multiple times)"
		// @Param speaker query string false "filter by UUID of any speaker (multiple times)"
		// @Param lang query string false "show description in given language"
		r.GET("/api/v1/recordings", func(c *gin.Context) {
			list := []*models.Recording{}
			db := ws.DB

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
				db = db.Joins("LEFT JOIN recording_tags ON recording_tags.recording_id = recordings.id").Where("tag_id IN (?)", ids)
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
				db = db.Joins("LEFT JOIN recording_speakers ON recording_speakers.recording_id = recordings.id").Where("speaker_id IN (?)", ids)
			}
			if err := db.Order("created_at DESC").Find(&list).Error; err != nil {
				c.JSON(http.StatusInternalServerError, web.HTTPError{
					Message: web.APIErrorInternalDatabase,
					Error:   err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, &list)
		})

		// @Summary Show Recording
		// @Description Show recording with all informations
		// @Produce  json
		// @Success 200 {object} models.Recording{formats=[]models.RecordingFormat}
		// @Failure 400 {object} HTTPError
		// @Failure 404 {object} HTTPError
		// @Router /api/v1/recording/{slug} [get]
		// @Param slug path string false "slug or uuid of recording"
		// @Param video_format query bool false "just format with video output"
		// @Param lang query string false "show description in given language"
		r.GET("/api/v1/recording/:slug", func(c *gin.Context) {
			slug := c.Params.ByName("slug")
			db := ws.DB.Joins("Channel").Joins("Event").Preload("Speakers")
			obj := models.Recording{}

			if str, ok := c.GetQuery("video_format"); ok {
				isVideo, err := strconv.ParseBool(str)
				if err != nil {
					c.JSON(http.StatusBadRequest, web.HTTPError{
						Message: web.APIErrorInvalidRequestFormat,
						Error:   err.Error(),
					})
					return
				}
				db = db.Preload("Formats", func(db *gorm.DB) *gorm.DB {
					return db.Where("is_video", isVideo).Order("quality ASC")
				})

			} else {
				db = db.Preload("Formats", func(db *gorm.DB) *gorm.DB {
					return db.Order("quality ASC")
				})
			}
			if str, ok := c.GetQuery("lang"); ok {
				db = db.Preload("Lang", func(db *gorm.DB) *gorm.DB {
					return db.Where("lang", str).Limit(1)
				}).Preload("Tags.Lang", func(db *gorm.DB) *gorm.DB {
					return db.Where("lang", str)
				})
			} else {
				db = db.Preload("Tags")
			}

			uuid, err := uuid.Parse(slug)
			if err != nil {
				db = db.Where("common_name", slug)
				obj.CommonName = slug
			} else {
				obj.ID = uuid
			}
			if err := db.First(&obj).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					c.JSON(http.StatusNotFound, web.HTTPError{
						Message: web.APIErrorNotFound,
						Error:   err.Error(),
					})
					return
				}
				c.JSON(http.StatusInternalServerError, web.HTTPError{
					Message: web.APIErrorInternalDatabase,
					Error:   err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, &obj)
		})
	})
}
