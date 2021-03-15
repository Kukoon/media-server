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

// @Summary Login status
// @Description show user_id and username if logged in
// @Accept json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 401 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/auth/status [get]
func init() {
	web.ModuleRegister(func(r *gin.Engine, ws *web.Service) {
		r.GET("/api/v1/auth/status", func(c *gin.Context) {
			session := sessions.Default(c)

			v := session.Get("user_id")
			if v == nil {
				c.JSON(http.StatusUnauthorized, web.HTTPError{
					Message: APIErrorNoSession,
				})
				return
			}

			id := uuid.MustParse(v.(string))

			d := &models.User{ID: id}
			if err := ws.DB.First(d).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					c.JSON(http.StatusUnauthorized, web.HTTPError{
						Message: APIErrorUserNotFound,
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

			c.JSON(http.StatusOK, d)
		})
	})
}
