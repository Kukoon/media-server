package auth

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/Kukoon/media-server/models"
	"github.com/Kukoon/media-server/web"
)

type login struct {
	Username string `json:"username" example:"kukoon"`
	Password string `json:"password" example:"super secret password"`
}

// @Summary Login
// @Description Login by username and password, you will get a cookie of current session
// @Accept json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/auth/login [post]
// @Param body body login false "login"
func init() {
	web.ModuleRegister(func(r *gin.Engine, ws *web.Service) {
		r.POST("/api/v1/auth/login", func(c *gin.Context) {
			var data login
			if err := c.BindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, web.HTTPError{
					Message: web.APIErrorInvalidRequestFormat,
					Error:   err.Error(),
				})
				return
			}

			d := &models.User{}
			if err := ws.DB.Where(map[string]interface{}{"username": data.Username}).First(d).Error; err != nil {
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
			if !d.ValidatePassword(data.Password) {
				c.JSON(http.StatusUnauthorized, web.HTTPError{
					Message: APIErrorIncorrectPassword,
				})
				return
			}

			session := sessions.Default(c)
			session.Set("user_id", d.ID.String())
			if err := session.Save(); err != nil {
				c.JSON(http.StatusBadRequest, web.HTTPError{
					Message: APIErrorCreateSession,
					Error:   err.Error(),
				})
				return
			}

			c.JSON(200, d)
		})
	})
}
