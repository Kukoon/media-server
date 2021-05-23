package stream

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/Kukoon/media-server/web"
	"github.com/Kukoon/media-server/web/metrics"
)

func init() {
	web.ModuleRegister(func(r *gin.Engine, ws *web.Service) {

		endpoints := make(map[uuid.UUID]*endpoint)

		r.GET("/ws/v1/stream/:uuid", func(c *gin.Context) {
			idString := c.Params.ByName("uuid")
			id, err := uuid.Parse(idString)
			if err != nil {
				return
			}
			// TODO validate
			s, ok := endpoints[id]
			if !ok {
				s = NewEndpoint()
				endpoints[id] = s
			}
			prometheus.MustRegister(prometheus.NewGaugeFunc(
				prometheus.GaugeOpts{
					Namespace:   metrics.NAMESPACE,
					Name:        "stream_clients",
					Help:        "count of current clients (count by websocket)",
					ConstLabels: prometheus.Labels{"stream": idString},
				},
				func() float64 {
					return float64(len(s.Subscribers))
				},
			))
			s.Handler(c)
		})
	})
}
