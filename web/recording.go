package web

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kukoon/media-server/models"
)

// @Summary List all Recordings
// @Description Show a list of all recordings
// @Produce  json
// @Success 200 {array} models.Recording
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /v1/recordings [get]
// @Param channel query string false "filter by UUID of a user"
// @Param lang query string false "show description in given language"
func (ws *Webservice) apiRecordingList(c *gin.Context) {
	list := []*models.Recording{}
	db := ws.DB
	if str, ok := c.GetQuery("channel"); ok {
		uuid, err := uuid.Parse(str)
		if err != nil {
			c.JSON(http.StatusBadRequest, HTTPError{
				Message: APIErrorInvalidRequestFormat,
				Error:   err.Error(),
			})
			return
		}
		db = db.Where("channel_id", uuid)
	}
	if str, ok := c.GetQuery("lang"); ok {
		db = db.Preload("RecordingLang", func(db *gorm.DB) *gorm.DB {
			return db.Where("lang", str)
		})
	}
	if err := db.Joins("Channel").Find(&list).Error; err != nil {
		c.JSON(http.StatusInternalServerError, HTTPError{
			Message: APIErrorInternalDatabase,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &list)
}

// @Summary Show Recording
// @Description Show recording with all informations
// @Produce  json
// @Success 200 {object} models.Recording{formats=[]models.RecordingFormat}
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Router /v1/recording/{slug} [get]
// @Param slug path string false "slug or uuid of recording"
// @Param video_format query bool false "just format with video output"
// @Param lang query string false "show description in given language"
func (ws *Webservice) apiRecordingGet(c *gin.Context) {
	slug := c.Params.ByName("slug")
	db := ws.DB.Joins("Channel")
	obj := models.Recording{}

	if str, ok := c.GetQuery("video_format"); ok {
		b, err := strconv.ParseBool(str)
		if err != nil {
			c.JSON(http.StatusBadRequest, HTTPError{
				Message: APIErrorInvalidRequestFormat,
				Error:   err.Error(),
			})
			return
		}
		db = db.Preload("Formats", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_video", b)
		})

	} else {
		db = db.Preload("Formats")
	}
	if str, ok := c.GetQuery("lang"); ok {
		db = db.Preload("RecordingLang", func(db *gorm.DB) *gorm.DB {
			return db.Where("lang", str).Limit(1)
		})
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
			c.JSON(http.StatusNotFound, HTTPError{
				Message: APIErrorNotFound,
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, HTTPError{
			Message: APIErrorInternalDatabase,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &obj)
}
