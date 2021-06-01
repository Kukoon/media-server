package channel

import (
	"net/http"
	"testing"

	"dev.sum7.eu/genofire/golang-lib/database"
	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/stretchr/testify/assert"

	"github.com/Kukoon/media-server/models"
)

func TestAPIChannelList(t *testing.T) {
	assert := assert.New(t)
	s := webtest.New(assert)
	assert.NotNil(s)
	s.DatabaseMigration(func(db *database.Database) {
		models.SetupMigration(db)
	})

	list := []*models.Channel{}
	// GET
	s.Request(http.MethodGet, "/api/v1/channels", nil, http.StatusOK, &list)
}
