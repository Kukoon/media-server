package models

import (
	"fmt"

	"dev.sum7.eu/genofire/golang-lib/database"
	"dev.sum7.eu/genofire/golang-lib/web/webtest"
)

func DatabaseForTesting() *database.Database {
	dbConfig := database.Database{
		Testdata: true,
		Debug:    false,
		LogLevel: 0,
	}
	dbConfig.Connection.URI = webtest.DBConnection
	SetupMigration(&dbConfig)
	err := dbConfig.ReRun()
	if err != nil {
		fmt.Println(err.Error())
	}
	return &dbConfig
}
