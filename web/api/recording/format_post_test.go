package recording

import (
	"net/http"
	"testing"

	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"dev.sum7.eu/genofire/golang-lib/web"

	"github.com/Kukoon/media-server/models"
)

func TestAPIFormatPost(t *testing.T) {
	assert := assert.New(t)
	s, err := webtest.NewWithDBSetup(bindTest, models.SetupMigration)
	assert.NoError(err)
	defer s.Close()
	assert.NotNil(s)

	hErr := web.HTTPError{}
	// GET - not found
	err = s.Request(http.MethodPost, "/api/v1/recording/"+models.TestRecording1ID.String()+"/format", nil, http.StatusUnauthorized, &hErr)
	assert.NoError(err)
	assert.Equal(auth.ErrAPINoSession.Error(), hErr.Message)

	err = s.Login(webtest.Login{
		Username: "kukoon",
		Password: "CHANGEME",
	})
	assert.NoError(err)

	hErr = web.HTTPError{}
	// GET - id
	err = s.Request(http.MethodPost, "/api/v1/recording/"+models.TestRecording1ID.String()+"/format", nil, http.StatusBadRequest, &hErr)
	assert.NoError(err)
	assert.Equal(web.ErrAPIInvalidRequestFormat.Error(), hErr.Message)

	req := models.RecordingFormat{
		URL: "de",
	}
	resp := models.RecordingFormat{}
	// GET - id
	err = s.Request(http.MethodPost, "/api/v1/recording/"+models.TestRecording1ID.String()+"/format", &req, http.StatusOK, &resp)
	assert.NoError(err)
	assert.NotEqual(uuid.Nil, resp.ID)

	err = s.DB.DB.Delete(&resp).Error
	assert.NoError(err)
}
