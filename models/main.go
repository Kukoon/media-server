package models

import (
	"dev.sum7.eu/genofire/golang-lib/database"
)

func SetupMigration(db *database.Database) {
	db.AddMigration(migrations...)
	db.AddMigrationTestdata(testdata...)
}
