package event

import (
	"net/http"
	"testing"

	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/stretchr/testify/assert"

	"dev.sum7.eu/genofire/golang-lib/web"

	"github.com/Kukoon/media-server/models"
)

func TestAPIPut(t *testing.T) {
	assert := assert.New(t)
	s, err := webtest.NewWithDBSetup(bindTest, models.SetupMigration)
	assert.NoError(err)
	defer s.Close()
	assert.NotNil(s)

	hErr := web.HTTPError{}
	// GET - not found
	err = s.Request(http.MethodPut, "/api/v1/event/"+models.TestEventID1.String(), nil, http.StatusUnauthorized, &hErr)
	assert.NoError(err)
	assert.Equal(auth.ErrAPINoSession.Error(), hErr.Message)

	err = s.Login(webtest.Login{
		Username: "kukoon",
		Password: "CHANGEME",
	})
	assert.NoError(err)

	hErr = web.HTTPError{}
	// GET - id
	err = s.Request(http.MethodPut, "/api/v1/event/00000000-0000-0000-0000-000000000001", nil, http.StatusNotFound, &hErr)
	assert.NoError(err)
	assert.Equal(web.ErrAPINotFound.Error(), hErr.Message)

	hErr = web.HTTPError{}
	// GET - id
	err = s.Request(http.MethodPut, "/api/v1/event/"+models.TestEventID1.String(), nil, http.StatusBadRequest, &hErr)
	assert.NoError(err)
	assert.Equal(web.ErrAPIInvalidRequestFormat.Error(), hErr.Message)

	req := models.Event{
		Name: "nope",
	}
	resp := models.Event{}
	// GET - id
	err = s.Request(http.MethodPut, "/api/v1/event/"+models.TestEventID1.String(), &req, http.StatusOK, &resp)
	assert.NoError(err)
	assert.Equal(models.TestEventID1, resp.ID)
	assert.Equal("nope", resp.Name)
}
