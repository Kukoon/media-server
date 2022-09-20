package channel

import (
	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"github.com/gin-gonic/gin"
)

func bindTest(r *gin.Engine, ws *web.Service) {
	Bind(r, ws, nil, &ConfigStream{
		Ingress: StreamIngress{
			WebRTC: "",
			RTMP:   "",
		},
		Streams: []StreamSource{
			{
				Label: "Live",
				Type:  "ll-hls",
				File:  "http://localhost/app/{ID}/ll-hls.m3u",
			},
		},
	})
	auth.Register(r, ws)
}
