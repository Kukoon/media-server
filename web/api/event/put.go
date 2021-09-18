package event

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

// @Summary Edit Event
// @Description Edit event by ID
// @Tags event
// @Produce  json
// @Success 200 {object} models.Event
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/event/{event_id} [put]
// @Param event_id path string false "uuid of event"
// @Param body body Event false "new values in event"
// @Security ApiKeyAuth
func apiPut(r *gin.Engine, ws *web.Service) {
	r.PUT("/api/v1/event/:uuid", auth.MiddlewarePermissionParamUUID(ws, models.Event{}), func(c *gin.Context) {
		old := models.Event{
			ID: uuid.MustParse(c.Params.ByName("uuid")),
		}
		if err := ws.DB.First(&old).Error; err != nil {
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
		var data models.Event
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, web.HTTPError{
				Message: web.ErrAPIInvalidRequestFormat.Error(),
				Error:   err.Error(),
			})
			return
		}

		data.ID = old.ID
		data.OwnerID = old.OwnerID

		if err := ws.DB.Omit("Owner.*").Save(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, web.HTTPError{
				Message: web.ErrAPIInternalDatabase.Error(),
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &data)
	})
}
