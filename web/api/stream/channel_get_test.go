package stream

import (
	"net/http"
	"testing"
	"time"

	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/stretchr/testify/assert"

	"dev.sum7.eu/genofire/golang-lib/web"

	"github.com/Kukoon/media-server/models"
)

func TestAPIChannelGet(t *testing.T) {
	assert := assert.New(t)
	s, err := webtest.NewWithDBSetup(apiChannelGet, models.SetupMigration)
	assert.NoError(err)
	defer s.Close()
	assert.NotNil(s)

	now := time.Now()
	err = s.DB.DB.Model(&models.Stream{}).Where("id", models.TestStreamID1).Update("end_at", now.Add(time.Hour)).Error
	assert.NoError(err)

	obj := models.PublicStream{}
	// GET - common name
	err = s.Request(http.MethodGet, "/api/v1/channel/kukoon/stream", nil, http.StatusOK, &obj)
	assert.NoError(err)

	obj = models.PublicStream{}
	// GET - id
	err = s.Request(http.MethodGet, "/api/v1/channel/df1555f5-7046-4f7a-adcc-195b73949723/stream?lang=de", nil, http.StatusOK, &obj)
	assert.NoError(err)

	hErr := web.HTTPError{}
	// GET - not found
	err = s.Request(http.MethodGet, "/api/v1/channel/00000000-0000-0000-0000-000000000001/stream", nil, http.StatusNotFound, &hErr)
	assert.NoError(err)
}
