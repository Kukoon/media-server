package main

import (
	"flag"

	"github.com/bdlm/log"

	"dev.sum7.eu/genofire/golang-lib/database"
	"dev.sum7.eu/genofire/golang-lib/file"
	"dev.sum7.eu/genofire/golang-lib/web"
	apiStatus "dev.sum7.eu/genofire/golang-lib/web/api/status"
	webM "dev.sum7.eu/genofire/golang-lib/web/metrics"
	ovenAPI "dev.sum7.eu/genofire/oven-exporter/api"

	"github.com/Kukoon/media-server/models"
	webOWN "github.com/Kukoon/media-server/web"
)

var VERSION = "development"

type configData struct {
	Database  database.Database `toml:"database"`
	Webserver web.Service       `toml:"webserver"`
	OvenAPI   ovenAPI.Client    `toml:"oven"`
	OvenVHost string            `toml:"oven_vhost"`
	OvenApp   string            `toml:"oven_app"`
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
	if err := file.ReadTOML(configPath, config); err != nil {
		log.Panicf("open config file: %s", err)
	}
	config.OvenAPI.SetToken(config.OvenAPI.Token)
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

	config.Webserver.ModuleRegister(webOWN.Bind(&config.OvenAPI, config.OvenVHost, config.OvenApp))

	if err := config.Webserver.Run(); err != nil {
		log.Fatal(err)
	}

}
