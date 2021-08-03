package stream

import (
	"errors"
	"net/http"
	"time"

	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kukoon/media-server/models"
)

// @Summary Show current Stream of channel
// @Description Show stream with all informations
// @Produce  json
// @Success 200 {object} models.PublicStream{}
// @Failure 400 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Router /api/v1/stream/{slug} [get]
// @Param slug path string false "slug or uuid of stream"
// @Param lang query string false "show description in given language"
func apiGet(r *gin.Engine, ws *web.Service) {
	r.GET("/api/v1/stream/:slug", func(c *gin.Context) {
		slug := c.Params.ByName("slug")
		db := ws.DB.Joins("Event").Preload("Speakers").Joins("Channel")
		if id, err := uuid.Parse(slug); err == nil {
			db = db.Where("streams.channel_id=?", id)
		} else {
			db = db.Where("\"Channel\".\"common_name\"=?", slug)
		}
		obj := models.Stream{}

		if str, ok := c.GetQuery("lang"); ok {
			db = db.Preload("Lang", func(db *gorm.DB) *gorm.DB {
				return db.Where("lang", str).Limit(1)
			}).Preload("Tags.Lang", func(db *gorm.DB) *gorm.DB {
				return db.Where("lang", str)
			})
		} else {
			db = db.Preload("Tags")
		}

		now := time.Now()

		db = db.
			Where("listen_at < ?", now).
			Where("start_at < ?", now).
			Order("start_at DESC")

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

		c.JSON(http.StatusOK, obj.GetPublic())
	})
}
