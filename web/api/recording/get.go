package recording

import (
	"errors"
	"net/http"
	"strconv"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kukoon/media-server/models"
)

// @Summary Show Recording
// @Description Show recording with all informations
// @Produce  json
// @Success 200 {object} models.Recording{formats=[]models.RecordingFormat}
// @Failure 400 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Router /api/v1/recording/{slug} [get]
// @Param slug path string false "slug or uuid of recording"
// @Param video_format query bool false "just format with video output"
// @Param count_viewer query bool false "count this request as an viewer"
// @Param lang query string false "show description in given language"
func apiGet(r *gin.Engine, ws *web.Service) {
	r.GET("/api/v1/recording/:uuid", func(c *gin.Context) {
		slug := c.Params.ByName("uuid")
		db := ws.DB.Joins("Channel").Joins("Event").Preload("Speakers")
		obj := models.Recording{}

		countViewer := false

		if str, ok := c.GetQuery("count_viewer"); ok {
			b, err := strconv.ParseBool(str)
			if err != nil {
				c.JSON(http.StatusBadRequest, web.HTTPError{
					Message: web.ErrAPIInvalidRequestFormat.Error(),
					Error:   err.Error(),
				})
				return
			}
			countViewer = b
		}
		if str, ok := c.GetQuery("video_format"); ok {
			isVideo, err := strconv.ParseBool(str)
			if err != nil {
				c.JSON(http.StatusBadRequest, web.HTTPError{
					Message: web.ErrAPIInvalidRequestFormat.Error(),
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
			db = db.Where("recordings.common_name", slug)
		} else {
			obj.ID = uuid
		}

		if err := db.First(&obj).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, web.HTTPError{
					Message: web.ErrAPINotFound.Error(),
					Error:   err.Error(),
				})
				return
			}
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}
		// have permission - own channel
		if !obj.Public {
			id, ok := auth.GetCurrentUserID(c)
			if !ok {
				c.JSON(http.StatusNotFound, web.HTTPError{
					Message: web.ErrAPINotFound.Error(),
				})
				return
			}
			if a, err := obj.HasPermission(ws.DB, id, obj.ID); err != nil || a == nil {
				c.JSON(http.StatusNotFound, web.HTTPError{
					Message: web.ErrAPINotFound.Error(),
				})
				return
			}
		}

		if countViewer {
			ws.DB.Model(&obj).Update("viewers", obj.Viewers+1)
		}

		c.JSON(http.StatusOK, &obj)
	})
}
