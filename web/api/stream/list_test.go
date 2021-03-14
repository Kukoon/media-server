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
