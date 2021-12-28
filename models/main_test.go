package models

import (
	"fmt"

	"dev.sum7.eu/genofire/golang-lib/database"
	"dev.sum7.eu/genofire/golang-lib/web/webtest"
)

func DatabaseForTesting() *database.Database {
	dbConfig := database.Database{
		Connection: webtest.DBConnection,
		Testdata:   true,
		Debug:      false,
		LogLevel:   0,
	}
	SetupMigration(&dbConfig)
	err := dbConfig.ReRun()
	if err != nil {
		fmt.Println(err.Error())
	}
	return &dbConfig
}
