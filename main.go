package main

import (
	"flag"
	"path/filepath"
	"strings"

	"dev.sum7.eu/genofire/golang-lib/database"
	"dev.sum7.eu/genofire/golang-lib/web"
	apiStatus "dev.sum7.eu/genofire/golang-lib/web/api/status"
	webM "dev.sum7.eu/genofire/golang-lib/web/metrics"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"go.uber.org/zap"

	"github.com/Kukoon/media-server/models"
	"github.com/Kukoon/media-server/oven"
	webOWN "github.com/Kukoon/media-server/web"
	"github.com/Kukoon/media-server/web/api/channel"
)

var VERSION = "development"

var configExtParser = map[string]koanf.Parser{
	".json": json.Parser(),
	".toml": toml.Parser(),
	".yaml": yaml.Parser(),
	".yml":  yaml.Parser(),
}

type configData struct {
	Log        *zap.Config          `config:"log"`
	Database   database.Database    `config:"database"`
	Webserver  web.Service          `config:"webserver"`
	Oven       oven.Service         `config:"oven"`
	StreamURLs channel.ConfigStream `config:"stream_urls"`
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

	k := koanf.New("/")

	if configPath != "" {
		fileExt := filepath.Ext(configPath)
		parser, ok := configExtParser[fileExt]
		if !ok {
			log.Panic("unsupported file extension:",
				zap.String("config-path", configPath),
				zap.String("file-ext", fileExt),
			)
		}
		if err := k.Load(file.Provider(configPath), parser); err != nil {
			log.Panic("load file config:", zap.Error(err))
		}
	}

	if err := k.Load(env.Provider("MEDIA_SERVER_", "/", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "MEDIA_SERVER_")), "__", "/", -1)
	}), nil); err != nil {
		log.Panic("load env:", zap.Error(err))
	}

	config := &configData{}
	if err := k.UnmarshalWithConf("", &config, koanf.UnmarshalConf{Tag: "config"}); err != nil {
		log.Panic("reading config", zap.Error(err))
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

	config.Webserver.ModuleRegister(webOWN.Bind(&config.Oven, &config.StreamURLs))

	config.Oven.Run(log)

	if err := config.Webserver.Run(log); err != nil {
		log.Fatal("crash webserver", zap.Error(err))
	}

}
