package stream

import (
	"net/http"
	"testing"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/stretchr/testify/assert"

	"github.com/Kukoon/media-server/models"
)

func TestAPILangList(t *testing.T) {
	assert := assert.New(t)
	s, err := webtest.NewWithDBSetup(bindTest, models.SetupMigration)
	assert.NoError(err)
	defer s.Close()
	assert.NotNil(s)

	resp := []*models.StreamLang{}
	// GET - id
	err = s.Request(http.MethodGet, "/api/v1/stream/"+models.TestStreamID1.String()+"/langs", nil, http.StatusOK, &resp)
	assert.NoError(err)
	assert.Len(resp, 1)
	if len(resp) > 1 {
		lang := resp[0]
		assert.Equal(models.TestStream1IDLang1, lang.ID)
	}

	hErr := web.HTTPError{}
	// GET - id
	err = s.Request(http.MethodGet, "/api/v1/stream/invalid/langs", nil, http.StatusBadRequest, &hErr)
	assert.NoError(err)
	assert.Equal(web.ErrAPIInvalidRequestFormat.Error(), hErr.Message)

	resp = []*models.StreamLang{}
	// GET - id
	err = s.Request(http.MethodGet, "/api/v1/stream/00000000-0000-0000-0000-000000000001/langs", nil, http.StatusOK, &resp)
	assert.NoError(err)
	assert.Len(resp, 0)

	resp = []*models.StreamLang{}
	// GET - id
	err = s.Request(http.MethodGet, "/api/v1/stream/"+models.TestStreamID1.String()+"/langs?lang=en", nil, http.StatusOK, &resp)
	assert.NoError(err)
	assert.Len(resp, 0)

	resp = []*models.StreamLang{}
	// GET - id
	err = s.Request(http.MethodGet, "/api/v1/stream/"+models.TestStreamID1.String()+"/langs?lang=de", nil, http.StatusOK, &resp)
	assert.NoError(err)
	assert.Len(resp, 1)
}
