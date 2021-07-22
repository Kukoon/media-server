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
	s, err := webtest.NewWithDBSetup(apiList, models.SetupMigration)
	assert.NoError(err)
	defer s.Close()
	assert.NotNil(s)

	list := []*models.PublicStream{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/streams", nil, http.StatusOK, &list)
	assert.NoError(err)
}
