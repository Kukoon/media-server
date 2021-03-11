package channel

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Kukoon/media-server/models"
	"github.com/Kukoon/media-server/web"
	"github.com/Kukoon/media-server/web/webtest"
)

func TestAPIChannelList(t *testing.T) {
	assert := assert.New(t)
	s := webtest.New(assert)
	assert.NotNil(s)

	list := []*models.Channel{}
	// GET
	s.Request(http.MethodGet, "/api/v1/channels", nil, http.StatusOK, &list)
}

func TestAPIChannelGet(t *testing.T) {
	assert := assert.New(t)
	s := webtest.New(assert)
	assert.NotNil(s)

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
