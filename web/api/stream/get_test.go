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

func TestAPIStreamGet(t *testing.T) {
	assert := assert.New(t)
	s := webtest.New(assert)
	assert.NotNil(s)
	s.DatabaseMigration(func(db *database.Database) {
		models.SetupMigration(db)
	})

	obj := models.PublicStream{}
	// GET - common name
	s.Request(http.MethodGet, "/api/v1/stream/kukoon", nil, http.StatusOK, &obj)

	obj = models.PublicStream{}
	// GET - id
	s.Request(http.MethodGet, "/api/v1/stream/df1555f5-7046-4f7a-adcc-195b73949723?lang=de", nil, http.StatusOK, &obj)

	/* TODO oO - that should not happen
	hErr := web.HTTPError{}
	// GET - not found
	s.Request(http.MethodGet, "/api/v1/stream/00000000-0000-0000-0000-000000000001", nil, http.StatusNotFound, &hErr)
	*/
}
