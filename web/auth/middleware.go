package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Kukoon/media-server/models"
	"github.com/Kukoon/media-server/web"
)

func MiddlewareLogin(ws *web.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := GetCurrentUserID(c)
		if !ok {
			c.Abort()
		}
	}
}

func MiddlewarePermissionParamUUID(ws *web.Service, obj models.HasPermission) gin.HandlerFunc {
	return MiddlewarePermissionParam(ws, obj, "uuid")
}
func MiddlewarePermissionParam(ws *web.Service, obj models.HasPermission, param string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := GetCurrentUserID(c)
		if !ok {
			c.Abort()
		}
		objID, err := uuid.Parse(c.Params.ByName(param))
		if err != nil {
			c.JSON(http.StatusUnauthorized, web.HTTPError{
				Message: web.APIErrorInvalidRequestFormat,
				Error:   err.Error(),
			})
			c.Abort()
		}
		_, err = obj.HasPermission(ws.DB, userID, objID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, web.HTTPError{
				Message: http.StatusText(http.StatusUnauthorized),
				Error:   err.Error(),
			})
			c.Abort()
		}
	}
}
