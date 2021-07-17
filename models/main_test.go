package models

import (
	"fmt"

	"dev.sum7.eu/genofire/golang-lib/database"
)

func DatabaseForTesting() *database.Database {
	dbConfig := database.Database{
		Connection: DBTestConnection,
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
