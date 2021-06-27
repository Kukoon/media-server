package stream

import (
	"net/http"
	"testing"

	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/stretchr/testify/assert"

	// "dev.sum7.eu/genofire/golang-lib/web"

	"github.com/Kukoon/media-server/models"
)

func TestAPIStreamList(t *testing.T) {
	assert := assert.New(t)
	s := webtest.New(assert)
	defer s.Close()
	assert.NotNil(s)
	models.SetupMigration(s.DB)

	list := []*models.PublicStream{}
	// GET
	s.Request(http.MethodGet, "/api/v1/streams", nil, http.StatusOK, &list)
}
