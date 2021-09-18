package speaker

import (
	"net/http"
	"testing"

	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/stretchr/testify/assert"

	"dev.sum7.eu/genofire/golang-lib/web"

	"github.com/Kukoon/media-server/models"
)

func TestAPIDelete(t *testing.T) {
	assert := assert.New(t)
	s, err := webtest.NewWithDBSetup(bindTest, models.SetupMigration)
	assert.NoError(err)
	defer s.Close()
	assert.NotNil(s)

	hErr := web.HTTPError{}
	// no permission
	err = s.Request(http.MethodDelete, "/api/v1/speaker/"+models.TestStream1IDSpeaker1.String(), nil, http.StatusUnauthorized, &hErr)
	assert.NoError(err)
	assert.Equal(auth.ErrAPINoSession.Error(), hErr.Message)

	err = s.Login(webtest.Login{
		Username: "kukoon",
		Password: "CHANGEME",
	})
	assert.NoError(err)

	hErr = web.HTTPError{}
	// not exists
	err = s.Request(http.MethodDelete, "/api/v1/speaker/00000000-0000-0000-0000-000000000001", nil, http.StatusNotFound, &hErr)
	assert.NoError(err)
	assert.Equal(web.ErrAPINotFound.Error(), hErr.Message)

	resp := false
	// success
	err = s.Request(http.MethodDelete, "/api/v1/speaker/"+models.TestStream1IDSpeaker1.String(), nil, http.StatusOK, &resp)
	assert.NoError(err)
	assert.True(resp)
}
