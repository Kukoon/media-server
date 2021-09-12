package stream

import (
	"net/http"
	"testing"

	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/stretchr/testify/assert"

	"dev.sum7.eu/genofire/golang-lib/web"

	"github.com/Kukoon/media-server/models"
)

func TestAPIPostToRecording(t *testing.T) {
	assert := assert.New(t)
	s, err := webtest.NewWithDBSetup(bindTest, models.SetupMigration)
	assert.NoError(err)
	defer s.Close()
	assert.NotNil(s)

	hErr := web.HTTPError{}
	// GET - not found
	err = s.Request(http.MethodPost, "/api/v1/stream/"+models.TestStreamID1.String()+"/to-recording", nil, http.StatusUnauthorized, &hErr)
	assert.NoError(err)
	assert.Equal(auth.ErrAPINoSession.Error(), hErr.Message)

	err = s.Login(webtest.Login{
		Username: "kukoon",
		Password: "CHANGEME",
	})
	assert.NoError(err)

	hErr = web.HTTPError{}
	// GET - id
	err = s.Request(http.MethodPost, "/api/v1/stream/"+models.TestStreamID1.String()+"/to-recording", nil, http.StatusInternalServerError, &hErr)
	assert.NoError(err)
	assert.Equal(web.ErrAPIInternalDatabase.Error(), hErr.Message)

	hErr = web.HTTPError{}
	// GET - not authorized
	err = s.Request(http.MethodPost, "/api/v1/stream/"+models.TestStreamID2.String()+"/to-recording", nil, http.StatusNotFound, &hErr)
	assert.NoError(err)
	assert.Equal(web.ErrAPINotFound.Error(), hErr.Message)

	err = s.Login(webtest.Login{
		Username: "c3woc",
		Password: "CHANGEME",
	})
	assert.NoError(err)

	resp := models.Recording{}
	// GET - not authorized
	err = s.Request(http.MethodPost, "/api/v1/stream/"+models.TestStreamID2.String()+"/to-recording", nil, http.StatusOK, &resp)
	assert.NoError(err)
	assert.Equal(models.TestStreamID2, resp.ID)

	err = s.DB.DB.Delete(&resp).Error
	assert.NoError(err)
}
