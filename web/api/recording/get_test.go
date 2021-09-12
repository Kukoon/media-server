package recording

import (
	"net/http"
	"testing"

	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/stretchr/testify/assert"

	"dev.sum7.eu/genofire/golang-lib/web"

	"github.com/Kukoon/media-server/models"
)

func TestAPIGet(t *testing.T) {
	assert := assert.New(t)
	s, err := webtest.NewWithDBSetup(bindTest, models.SetupMigration)
	assert.NoError(err)
	defer s.Close()
	assert.NotNil(s)

	hErr := web.HTTPError{}
	// GET - not found
	err = s.Request(http.MethodGet, "/api/v1/recording/"+models.TestRecording19ID.String(), nil, http.StatusUnauthorized, &hErr)
	assert.NoError(err)
	assert.Equal(auth.ErrAPINoSession.Error(), hErr.Message)

	err = s.Login(webtest.Login{
		Username: "c3woc",
		Password: "CHANGEME",
	})
	assert.NoError(err)

	hErr = web.HTTPError{}
	// GET - not found
	err = s.Request(http.MethodGet, "/api/v1/recording/"+models.TestRecording19ID.String(), nil, http.StatusNotFound, &hErr)
	assert.NoError(err)
	assert.Equal(web.ErrAPINotFound.Error(), hErr.Message)

	err = s.Login(webtest.Login{
		Username: "kukoon",
		Password: "CHANGEME",
	})
	assert.NoError(err)

	resp := models.Recording{}
	// GET - id
	err = s.Request(http.MethodGet, "/api/v1/recording/"+models.TestRecording19ID.String(), nil, http.StatusOK, &resp)
	assert.NoError(err)
	assert.Equal(models.TestRecording19ID, resp.ID)
	assert.Nil(resp.Lang)
	viewers := resp.Viewers

	resp = models.Recording{}
	// GET - id
	err = s.Request(http.MethodGet, "/api/v1/recording/"+models.TestRecording19ID.String()+"?lang=de", nil, http.StatusOK, &resp)
	assert.NoError(err)
	assert.Equal(models.TestRecording19ID, resp.ID)
	assert.NotNil(resp.Lang)

	hErr = web.HTTPError{}
	// GET - id
	err = s.Request(http.MethodGet, "/api/v1/recording/"+models.TestRecording1ID.String()+"?count_viewer=de", nil, http.StatusBadRequest, &hErr)
	assert.NoError(err)
	assert.Equal(web.ErrAPIInvalidRequestFormat.Error(), hErr.Message)

	resp = models.Recording{}
	// GET - id
	err = s.Request(http.MethodGet, "/api/v1/recording/"+models.TestRecording1ID.String()+"?count_viewer=true", nil, http.StatusOK, &resp)
	assert.NoError(err)
	assert.Equal(models.TestRecording1ID, resp.ID)
	assert.Greater(resp.Viewers, viewers)
	assert.Len(resp.Formats, 4)

	hErr = web.HTTPError{}
	// GET - id
	err = s.Request(http.MethodGet, "/api/v1/recording/"+models.TestRecording1ID.String()+"?video_format=de", nil, http.StatusBadRequest, &hErr)
	assert.NoError(err)
	assert.Equal(web.ErrAPIInvalidRequestFormat.Error(), hErr.Message)

	resp = models.Recording{}
	// GET - id
	err = s.Request(http.MethodGet, "/api/v1/recording/"+models.TestRecording1ID.String()+"?video_format=true", nil, http.StatusOK, &resp)
	assert.NoError(err)
	assert.Equal(models.TestRecording1ID, resp.ID)
	assert.Len(resp.Formats, 3)

	hErr = web.HTTPError{}
	// GET - not found
	err = s.Request(http.MethodGet, "/api/v1/recording/00000000-0000-0000-0000-000000000001", nil, http.StatusNotFound, &hErr)
	assert.NoError(err)
	assert.Equal(web.ErrAPINotFound.Error(), hErr.Message)

	hErr = web.HTTPError{}
	// GET - id
	err = s.Request(http.MethodGet, "/api/v1/recording/de", nil, http.StatusNotFound, &hErr)
	assert.NoError(err)
	assert.Equal(web.ErrAPINotFound.Error(), hErr.Message)

	resp = models.Recording{}
	// GET - id
	err = s.Request(http.MethodGet, "/api/v1/recording/2020-12-polizeigewalt", nil, http.StatusOK, &resp)
	assert.NoError(err)
	assert.Equal(models.TestRecording1ID, resp.ID)
}
