package status

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Kukoon/media-server/web"
)

var (
	VERSION string      = ""
	UP      func() bool = func() bool {
		return true
	}
	EXTRAS interface{} = nil
)

type Status struct {
	Version string      `json:"version"`
	Up      bool        `json:"up"`
	Extras  interface{} `json:"extras,omitempty"`
}

// @Summary Show Status of current API
// @Description Show version and status
// @Produce  json
// @Success 200 {object} Status
// @Failure 400 {object} web.HTTPError
// @Failure 404 {object} web.HTTPError
// @Router /api/status [get]
func init() {
	web.ModuleRegister(func(r *gin.Engine, ws *web.Service) {
		r.GET("/api/status", func(c *gin.Context) {
			status := &Status{
				Version: VERSION,
				Up:      UP(),
				Extras:  EXTRAS,
			}
			if !status.Up {
				c.JSON(http.StatusInternalServerError, status)
				return
			}
			c.JSON(http.StatusOK, status)
		})
	})
}
