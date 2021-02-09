package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kukoon/media-server/models"
)

// @Summary List all Recordings
// @Description Show a list of all recordings
// @Produce  json
// @Success 200 {array} models.Recording
// @Failure 500 {object} string
// @Router /v1/recordings [get]
func (ws *Webservice) apiRecordingList(c *gin.Context) {
	list := []*models.Recording{}
	if err := ws.DB.Find(&list).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &list)
}

// @Summary Show Recording
// @Description Show recording with all informations
// @Produce  json
// @Success 200 {object} models.Recording{Formats=[]models.RecordingFormat}
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /v1/recording/{slug} [get]
// @Param slug path string false "slug or uuid of recording"
func (ws *Webservice) apiRecordingGet(c *gin.Context) {
	slug := c.Params.ByName("slug")
	db := ws.DB.Preload("Formats")
	obj := models.Recording{}

	uuid, err := uuid.Parse(slug)
	if err != nil {
		db = db.Where("common_name", slug)
		obj.CommonName = slug
	} else {
		obj.ID = uuid
	}
	if err := db.First(&obj).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &obj)
}
