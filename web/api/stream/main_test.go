package stream

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Kukoon/media-server/models"
	// "github.com/Kukoon/media-server/web"
	"github.com/Kukoon/media-server/web/webtest"
)

func TestAPIStreamList(t *testing.T) {
	assert := assert.New(t)
	s := webtest.New(assert)
	assert.NotNil(s)

	list := []*models.PublicStream{}
	// GET
	s.Request(http.MethodGet, "/api/v1/streams", nil, http.StatusOK, &list)
}

func TestAPIStreamGet(t *testing.T) {
	assert := assert.New(t)
	s := webtest.New(assert)
	assert.NotNil(s)

	obj := models.PublicStream{}
	// GET - common name
	s.Request(http.MethodGet, "/api/v1/stream/kukoon", nil, http.StatusOK, &obj)

	obj = models.PublicStream{}
	// GET - id
	s.Request(http.MethodGet, "/api/v1/stream/df1555f5-7046-4f7a-adcc-195b73949723", nil, http.StatusOK, &obj)

	/* TODO oO - that should not happen
	hErr := web.HTTPError{}
	// GET - not found
	s.Request(http.MethodGet, "/api/v1/stream/00000000-0000-0000-0000-000000000001", nil, http.StatusNotFound, &hErr)
	*/
}
