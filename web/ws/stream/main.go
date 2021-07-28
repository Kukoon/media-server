package stream

import (
	"time"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/metrics"
	"dev.sum7.eu/genofire/golang-lib/worker"
	"github.com/bdlm/log"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/Kukoon/media-server/oven"
)

// Bind to webservice
func Bind(oven *oven.Service) web.ModuleRegisterFunc {

	endpoints := make(map[uuid.UUID]*endpoint)

	w := worker.NewWorker(5*time.Second, func() {
		// check status of stream server
		resp, err := oven.Client.RequestDefaultListStreams()
		if err != nil {
			log.WithField("error", err).Warn("websocket status check for oven stream server")
		} else {
			running := make(map[string]interface{})
			for _, stream := range resp.Data {
				running[stream] = true
			}
			for stream, e := range endpoints {
				_, ok := running[stream.String()]
				e.Running = ok
			}
		}
	})
	w.Start()
	return func(r *gin.Engine, ws *web.Service) {

		r.GET("/ws/v1/stream/:uuid", func(c *gin.Context) {
			idString := c.Params.ByName("uuid")
			id, err := uuid.Parse(idString)
			if err != nil {
				return
			}
			// TODO validate
			s, ok := endpoints[id]
			if !ok {
				s = NewEndpoint(ws, id)
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
	}
}
