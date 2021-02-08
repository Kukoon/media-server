package main

import (
	"github.com/bdlm/log"

	"github.com/Kukoon/media-server/web"
)

func main() {
	ws := web.Webservice{
		Listen: ":8090",
	}
	if err := ws.Run(); err != nil {
		log.Fatal(err)
	}

}
