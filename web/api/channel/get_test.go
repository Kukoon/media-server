package channel

import (
	"net/http"
	"testing"

	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/stretchr/testify/assert"

	"dev.sum7.eu/genofire/golang-lib/web"

	"github.com/Kukoon/media-server/models"
)

func TestAPIGet(t *testing.T) {
	assert := assert.New(t)
	s, err := webtest.NewWithDBSetup(bindTest, models.SetupMigration)
	assert.NoError(err)
	defer s.Close()
	assert.NotNil(s)

	obj := Channel{}
	// GET - common name
	err = s.Request(http.MethodGet, "/api/v1/channel/kukoon", nil, http.StatusOK, &obj)
	assert.NoError(err)

	obj = Channel{}
	// GET - id
	err = s.Request(http.MethodGet, "/api/v1/channel/df1555f5-7046-4f7a-adcc-195b73949723", nil, http.StatusOK, &obj)
	assert.NoError(err)

	hErr := web.HTTPError{}
	// GET - not found
	err = s.Request(http.MethodGet, "/api/v1/channel/00000000-0000-0000-0000-000000000001", nil, http.StatusNotFound, &hErr)
	assert.NoError(err)
}
