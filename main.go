package main

import (
	"flag"

	"github.com/bdlm/log"

	"github.com/Kukoon/media-server/models"
	"github.com/Kukoon/media-server/runtime"
	"github.com/Kukoon/media-server/web"
)

type configData struct {
	Database  models.Database `toml:"database"`
	Webserver web.Webservice  `toml:"webserver"`
}

func main() {
	configPath := "config.toml"
	flag.StringVar(&configPath, "c", configPath, "path to configuration file")
	flag.Parse()
	config := &configData{}
	if err := runtime.ReadTOML(configPath, config); err != nil {
		log.Panicf("open config file: %s", err)
	}

	if err := config.Database.Run(); err != nil {
		log.Fatal(err)
	}

	config.Webserver.DB = config.Database.DB

	if err := config.Webserver.Run(); err != nil {
		log.Fatal(err)
	}

}
