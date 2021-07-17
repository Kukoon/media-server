package recording

import (
	"net/http"
	"testing"

	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/stretchr/testify/assert"

	// "dev.sum7.eu/genofire/golang-lib/web"

	"github.com/Kukoon/media-server/models"
)

func TestAPIRecordingsList(t *testing.T) {
	assert := assert.New(t)
	s, err := webtest.NewWithDBSetup(models.SetupMigration)
	assert.NoError(err)
	defer s.Close()
	assert.NotNil(s)

	list := []*models.Recording{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/recordings", nil, http.StatusOK, &list)
	assert.NoError(err)
}
