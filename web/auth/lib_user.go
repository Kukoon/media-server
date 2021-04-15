package auth

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kukoon/media-server/models"
	"github.com/Kukoon/media-server/web"
)

func GetCurrentUserID(c *gin.Context) (uuid.UUID, bool) {
	session := sessions.Default(c)

	v := session.Get("user_id")
	if v == nil {
		c.JSON(http.StatusUnauthorized, web.HTTPError{
			Message: APIErrorNoSession,
		})
		return uuid.Nil, false
	}

	id := uuid.MustParse(v.(string))
	return id, true
}

func GetCurrentUser(c *gin.Context, ws *web.Service) (*models.User, bool) {
	id, ok := GetCurrentUserID(c)
	if !ok {
		return nil, false
	}
	d := &models.User{ID: id}
	if err := ws.DB.First(d).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, web.HTTPError{
				Message: APIErrorUserNotFound,
				Error:   err.Error(),
			})
			return nil, false
		}
		c.JSON(http.StatusInternalServerError, web.HTTPError{
			Message: web.APIErrorInternalDatabase,
			Error:   err.Error(),
		})
		return nil, false
	}
	return d, true
}
