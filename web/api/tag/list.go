package tag

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/Kukoon/media-server/models"
)

// @Summary List all Tags
// @Description Show a list of all tags
// @Produce  json
// @Success 200 {array} models.Tag
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/tags [get]
func apiList(r *gin.Engine, ws *web.Service) {
	r.GET("/api/v1/tags", func(c *gin.Context) {
		list := []*models.Tag{}
		db := ws.DB
		if str, ok := c.GetQuery("lang"); ok {
			db = db.Preload("Lang", func(db *gorm.DB) *gorm.DB {
				return db.Where("lang", str)
			})
		}
		if err := db.Find(&list).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &list)
	})
}
