package channel

import (
	"net/http"
	"testing"

	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/stretchr/testify/assert"

	"github.com/Kukoon/media-server/models"
)

func TestAPIChannelList(t *testing.T) {
	assert := assert.New(t)
	s := webtest.New(assert)
	defer s.Close()
	assert.NotNil(s)
	models.SetupMigration(s.DB)

	list := []*models.Channel{}
	// GET
	s.Request(http.MethodGet, "/api/v1/channels", nil, http.StatusOK, &list)
}
