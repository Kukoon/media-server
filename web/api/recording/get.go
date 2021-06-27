package recording

import (
	"errors"
	"net/http"
	"strconv"

	"dev.sum7.eu/genofire/golang-lib/web"
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
// @Param lang query string false "show description in given language"
func init() {
	web.ModuleRegister(func(r *gin.Engine, ws *web.Service) {
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

			// TODO login - own channel
			db = db.Where("public", true)

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
