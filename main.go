package main

import (
	"flag"

	"dev.sum7.eu/genofire/golang-lib/database"
	"dev.sum7.eu/genofire/golang-lib/file"
	"dev.sum7.eu/genofire/golang-lib/web"
	apiStatus "dev.sum7.eu/genofire/golang-lib/web/api/status"
	webM "dev.sum7.eu/genofire/golang-lib/web/metrics"
	"go.uber.org/zap"

	"github.com/Kukoon/media-server/models"
	"github.com/Kukoon/media-server/oven"
	webOWN "github.com/Kukoon/media-server/web"
)

var VERSION = "development"

type configData struct {
	Log       *zap.Config       `toml:"log"`
	Database  database.Database `toml:"database"`
	Webserver web.Service       `toml:"webserver"`
	Oven      oven.Service      `toml:"oven"`
}

func main() {
	webM.VERSION = VERSION
	webM.NAMESPACE = "media_server"

	configPath := "config.toml"
	showVersion := false

	log, _ := zap.NewProduction()

	flag.StringVar(&configPath, "c", configPath, "path to configuration file")
	flag.BoolVar(&showVersion, "version", showVersion, "show current version")

	flag.Parse()

	if showVersion {
		log.Info("Version", zap.String("version", VERSION))
		return
	}

	config := &configData{}
	if err := file.ReadTOML(configPath, config); err != nil {
		log.Panic("open config file", zap.Error(err))
	}

	if config.Log != nil {
		l, err := config.Log.Build()
		if err != nil {
			log.Panic("generate logger from config", zap.Error(err))
		}
		log = l
	}
	config.Oven.Client.SetToken(config.Oven.Client.Token)
	models.SetupMigration(&config.Database)

	if err := config.Database.Run(); err != nil {
		log.Fatal("database setup", zap.Error(err))
	}

	config.Webserver.DB = config.Database.DB
	config.Oven.DB = config.Database.DB

	webM.UP = func() bool {
		return config.Database.Status() == nil
	}
	apiStatus.VERSION = webM.VERSION
	apiStatus.UP = webM.UP

	config.Webserver.ModuleRegister(webOWN.Bind(&config.Oven))

	config.Oven.Run(log)

	if err := config.Webserver.Run(log); err != nil {
		log.Fatal("crash webserver", zap.Error(err))
	}

}
