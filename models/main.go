package models

import (
	"github.com/Kukoon/media-server/database"
)

func SetupMigration(db *database.Database) {
	db.AddMigration(migrations...)
	db.AddMigrationTestdata(testdata...)
}
