package models

import (
	"dev.sum7.eu/genofire/golang-lib/database"
	gormigrate "github.com/genofire/gormigrate/v2"
)

var (
	// DBTestConnection - url to database on setting up default modul test
	DBTestConnection = "user=root password=root dbname=media_server host=localhost port=26257 sslmode=disable"
)

var migrations = []*gormigrate.Migration{}
var testdata = []*gormigrate.Migration{}

// SetupMigration - setup all data for migration
func SetupMigration(db *database.Database) {
	db.AddMigration(migrations...)
	db.AddMigrationTestdata(testdata...)
}
