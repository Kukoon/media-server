package tag

import (
	"net/http"
	"testing"

	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/stretchr/testify/assert"

	"github.com/Kukoon/media-server/models"
)

func TestAPITagList(t *testing.T) {
	assert := assert.New(t)
	s, err := webtest.NewWithDBSetup(apiList, models.SetupMigration)
	assert.NoError(err)
	defer s.Close()
	assert.NotNil(s)

	list := []*models.Tag{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/tags?lang=de", nil, http.StatusOK, &list)
	assert.NoError(err)
}
