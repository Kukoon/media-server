package auth

import (
	"net/http"

	"github.com/bdlm/log"
	"github.com/gin-gonic/gin"

	"github.com/Kukoon/media-server/web"
)

// @Summary Change Password
// @Description Change Password of current login user
// @Accept json
// @Produce  json
// @Success 200 {object} boolean "if password was saved (e.g. `true`)"
// @Failure 400 {object} web.HTTPError
// @Failure 401 {object} web.HTTPError
// @Failure 500 {object} web.HTTPError
// @Router /api/v1/my/auth/password [post]
// @Security ApiKeyAuth
// @Param body body string false "new password"
func init() {
	web.ModuleRegister(func(r *gin.Engine, ws *web.Service) {
		r.POST("/api/v1/my/auth/password", MiddlewareLogin(ws), func(c *gin.Context) {
			d, ok := GetCurrentUser(c, ws)
			if !ok {
				return
			}
			var password string
			if err := c.BindJSON(&password); err != nil {
				c.JSON(http.StatusBadRequest, web.HTTPError{
					Message: web.APIErrorInvalidRequestFormat,
					Error:   err.Error(),
				})
				return
			}
			if err := d.SetPassword(password); err != nil {
				c.JSON(http.StatusInternalServerError, web.HTTPError{
					Message: APIErrroCreatePassword,
					Error:   err.Error(),
				})
				return
			}

			result := ws.DB.Save(&d)
			if err := result.Error; err != nil {
				c.JSON(http.StatusInternalServerError, web.HTTPError{
					Message: web.APIErrorInternalDatabase,
					Error:   err.Error(),
				})
				return
			}
			if result.RowsAffected > 1 {
				log.Panicf("there should not be more then 1 user with the same email, it was %d session", result.RowsAffected)
			}

			c.JSON(http.StatusOK, result.RowsAffected == 1)
		})
	})
}
