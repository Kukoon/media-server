package stream

import (
	"errors"
	"net/http"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kukoon/media-server/models"
)

// @Summary Clone Stream to Recording
// @Description Clone Stream to Recording by ID
// @Tags stream,recording
// @Produce  json
// @Success 200 {object} models.Recording
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/stream/{stream_id} [post]
// @Param stream_id path string false "uuid of stream"
// @Security ApiKeyAuth
func apiPostToRecording(r *gin.Engine, ws *web.Service) {
	r.POST("/api/v1/stream/:uuid/to-recording", auth.MiddlewarePermissionParamUUID(ws, models.Stream{}), func(c *gin.Context) {
		data := models.Stream{
			ID: uuid.MustParse(c.Params.ByName("uuid")),
		}

		if err := ws.DB.Preload("Speakers").Preload("Tags").First(&data).Error; err != nil {
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
		recording := models.Recording{
			ID:         data.ID,
			CreatedAt:  data.StartAt,
			ChannelID:  data.ChannelID,
			CommonName: data.CommonName,
			Poster:     data.Poster,
			Preview:    data.Preview,
			Public:     false,
			Listed:     false,
			Viewers:    0,
			EventID:    data.EventID,
			Tags:       data.Tags,
			Speakers:   data.Speakers,
		}
		if err := ws.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&recording).Error; err != nil {
				return err
			}
			list := []*models.StreamLang{}
			if err := tx.Find(&list).Error; err != nil {
				return err
			}
			for _, lang := range list {
				if err := tx.Create(&models.RecordingLang{
					RecordingID: recording.ID,
					Lang:        lang.Lang,
					Title:       lang.Title,
					Subtitle:    lang.Subtitle,
					Short:       lang.Short,
					Long:        lang.Long,
				}).Error; err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &recording)
	})
}
