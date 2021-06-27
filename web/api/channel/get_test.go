package channel

import (
	"net/http"
	"testing"

	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/stretchr/testify/assert"

	"dev.sum7.eu/genofire/golang-lib/web"

	"github.com/Kukoon/media-server/models"
)

func TestAPIChannelGet(t *testing.T) {
	assert := assert.New(t)
	s := webtest.New(assert)
	defer s.Close()
	assert.NotNil(s)
	models.SetupMigration(s.DB)

	obj := models.Channel{}
	// GET - common name
	s.Request(http.MethodGet, "/api/v1/channel/kukoon", nil, http.StatusOK, &obj)

	obj = models.Channel{}
	// GET - id
	s.Request(http.MethodGet, "/api/v1/channel/df1555f5-7046-4f7a-adcc-195b73949723", nil, http.StatusOK, &obj)

	hErr := web.HTTPError{}
	// GET - not found
	s.Request(http.MethodGet, "/api/v1/channel/00000000-0000-0000-0000-000000000001", nil, http.StatusNotFound, &hErr)
}
