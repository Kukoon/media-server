package channel

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Kukoon/media-server/models"
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
