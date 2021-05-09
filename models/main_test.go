package models

import (
	"fmt"
)

func DatabaseForTesting() *Database {
	dbConfig := Database{
		Connection: "user=root password=root dbname=media_server host=localhost port=26257 sslmode=disable",
		Testdata:   true,
		Debug:      false,
		LogLevel:   0,
	}
	err := dbConfig.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	return &dbConfig
}
