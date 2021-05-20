package main

import (
	"flag"

	"github.com/bdlm/log"

	"github.com/Kukoon/media-server/database"
	"github.com/Kukoon/media-server/models"
	"github.com/Kukoon/media-server/runtime"
	"github.com/Kukoon/media-server/web"
	_ "github.com/Kukoon/media-server/web/all"
	apiStatus "github.com/Kukoon/media-server/web/api/status"
	webM "github.com/Kukoon/media-server/web/metrics"
)

var VERSION = "development"

type configData struct {
	Database  database.Database `toml:"database"`
	Webserver web.Service       `toml:"webserver"`
}

func main() {
	webM.VERSION = VERSION
	webM.NAMESPACE = "media_server"

	configPath := "config.toml"
	showVersion := false

	flag.StringVar(&configPath, "c", configPath, "path to configuration file")
	flag.BoolVar(&showVersion, "version", showVersion, "show current version")

	flag.Parse()

	if showVersion {
		log.WithField("version", VERSION).Info("Version")
		return
	}

	config := &configData{}
	if err := runtime.ReadTOML(configPath, config); err != nil {
		log.Panicf("open config file: %s", err)
	}

	models.SetupMigration(&config.Database)

	if err := config.Database.Run(); err != nil {
		log.Fatal(err)
	}

	config.Webserver.DB = config.Database.DB

	webM.UP = func() bool {
		return config.Database.Status() == nil
	}
	apiStatus.VERSION = webM.VERSION
	apiStatus.UP = webM.UP

	if err := config.Webserver.Run(); err != nil {
		log.Fatal(err)
	}

}
