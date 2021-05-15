package metrics

import (
	"strings"

	"github.com/gin-gonic/gin"

	// gin-prometheus
	"github.com/chenjiandongx/ginprom"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// db metrics
	gormPrometheus "gorm.io/plugin/prometheus"

	"github.com/Kukoon/media-server/web"
)

var (
	NAMESPACE string      = "service"
	VERSION   string      = ""
	UP        func() bool = func() bool {
		return true
	}
)

func init() {
	web.ModuleRegister(func(r *gin.Engine, ws *web.Service) {
		r.Use(ginprom.PromMiddleware(&ginprom.PromOpts{
			EndpointLabelMappingFn: func(c *gin.Context) string {
				url := c.Request.URL.Path
				for _, p := range c.Params {
					url = strings.Replace(url, p.Value, ":"+p.Key, 1)
				}
				return url
			},
		}))

		prometheus.MustRegister(prometheus.NewGaugeFunc(
			prometheus.GaugeOpts{
				Namespace:   NAMESPACE,
				Name:        "up",
				Help:        "is current version of service running",
				ConstLabels: prometheus.Labels{"version": VERSION},
			},
			func() float64 {
				if UP() {
					return 1
				}
				return 0
			},
		))

		if ws.DB != nil {
			ws.DB.Use(gormPrometheus.New(gormPrometheus.Config{
				DBName:          NAMESPACE,
				RefreshInterval: 15,
			}))

		}

		r.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
	})
}
