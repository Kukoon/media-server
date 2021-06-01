package stream

import (
	"net/http"
	"testing"

	"dev.sum7.eu/genofire/golang-lib/database"
	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/stretchr/testify/assert"

	// "dev.sum7.eu/genofire/golang-lib/web"

	"github.com/Kukoon/media-server/models"
)

func TestAPIStreamList(t *testing.T) {
	assert := assert.New(t)
	s := webtest.New(assert)
	assert.NotNil(s)
	s.DatabaseMigration(func(db *database.Database) {
		models.SetupMigration(db)
	})

	list := []*models.PublicStream{}
	// GET
	s.Request(http.MethodGet, "/api/v1/streams", nil, http.StatusOK, &list)
}
