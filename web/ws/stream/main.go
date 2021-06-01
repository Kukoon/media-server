package stream

import (
	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/metrics"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
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
				prometheus.MustRegister(prometheus.NewGaugeFunc(
					prometheus.GaugeOpts{
						Namespace:   metrics.NAMESPACE,
						Name:        "stream_viewers",
						Help:        "count of current viewers (count by websocket)",
						ConstLabels: prometheus.Labels{"stream": idString},
					},
					func() float64 {
						return float64(len(s.Subscribers))
					},
				))
				prometheus.MustRegister(prometheus.NewGaugeFunc(
					prometheus.GaugeOpts{
						Namespace:   metrics.NAMESPACE,
						Name:        "stream_chatusers",
						Help:        "count of current chat users (count by websocket)",
						ConstLabels: prometheus.Labels{"stream": idString},
					},
					func() float64 {
						return float64(len(s.usernames))
					},
				))
				endpoints[id] = s
			}
			s.Handler(c)
		})
	})
}
