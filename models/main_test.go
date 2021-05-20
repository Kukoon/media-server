package models

import (
	"fmt"

	"github.com/Kukoon/media-server/database"
)

func DatabaseForTesting() *database.Database {
	dbConfig := database.Database{
		Connection: "user=root password=root dbname=media_server host=localhost port=26257 sslmode=disable",
		Testdata:   true,
		Debug:      false,
		LogLevel:   0,
	}
	SetupMigration(&dbConfig)
	err := dbConfig.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	return &dbConfig
}
