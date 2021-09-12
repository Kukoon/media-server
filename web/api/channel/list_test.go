package channel

import (
	"net/http"
	"testing"

	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/stretchr/testify/assert"

	"github.com/Kukoon/media-server/models"
)

func TestAPIList(t *testing.T) {
	assert := assert.New(t)
	s, err := webtest.NewWithDBSetup(apiList, models.SetupMigration)
	assert.NoError(err)
	defer s.Close()
	assert.NotNil(s)

	list := []*models.Channel{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/channels", nil, http.StatusOK, &list)
	assert.NoError(err)
}
