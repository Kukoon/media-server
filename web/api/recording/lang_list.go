package recording

import (
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Kukoon/media-server/models"
)

// @Summary List all Language Description of an Recording
// @Description List all Descriptions/Languages of recording by ID
// @Tags recording
// @Produce  json
// @Success 200 {array} models.RecordingLang
// @Failure 400 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/recording/{recording_id}/langs [get]
// @Param recording_id path string false "uuid of recording"
// @Param lang query string false "show description in given language"
func apiLangList(r *gin.Engine, ws *web.Service) {
	r.GET("/api/v1/recording/:uuid/langs", func(c *gin.Context) {
		id, err := uuid.Parse(c.Params.ByName("uuid"))
		if err != nil {
			c.JSON(http.StatusBadRequest, web.HTTPError{
				Message: web.ErrAPIInvalidRequestFormat.Error(),
				Error:   err.Error(),
			})
			return
		}
		db := ws.DB.Where("recording_id = ?", id)
		if str, ok := c.GetQuery("lang"); ok {
			db = db.Where("lang", str)
		}

		list := []*models.RecordingLang{}
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
