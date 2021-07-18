package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"

	"github.com/bdlm/log"
	"github.com/gin-gonic/gin"

	"dev.sum7.eu/genofire/golang-lib/database"
	"dev.sum7.eu/genofire/golang-lib/file"
	"dev.sum7.eu/genofire/golang-lib/web"
	apiStatus "dev.sum7.eu/genofire/golang-lib/web/api/status"
	webM "dev.sum7.eu/genofire/golang-lib/web/metrics"
	"github.com/Kukoon/media-server/models"
	_ "github.com/Kukoon/media-server/web"
)

var VERSION = "development"

type configData struct {
	Database  database.Database `toml:"database"`
	Webserver web.Service       `toml:"webserver"`
}

//go:embed templates/*
var templateFS embed.FS

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
	web.ModuleRegister(func(r *gin.Engine, ws *web.Service) {
		templ := template.Must(template.New("").Funcs(r.FuncMap).ParseFS(templateFS, "templates/partials/*"))
		templ = template.Must(templ.ParseFS(templateFS, "templates/views/**/*"))
		fmt.Println(templ.DefinedTemplates())
		r.SetHTMLTemplate(templ)
	})
	if err := config.Webserver.Run(); err != nil {
		log.Fatal(err)
	}

}
