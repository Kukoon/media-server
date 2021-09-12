package channel

import (
	"net/http"
	"testing"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/stretchr/testify/assert"

	"github.com/Kukoon/media-server/models"
)

func TestAPIListMy(t *testing.T) {
	assert := assert.New(t)
	s, err := webtest.NewWithDBSetup(bindTest, models.SetupMigration)
	assert.NoError(err)
	defer s.Close()
	assert.NotNil(s)

	hErr := &web.HTTPError{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/my/channels", nil, http.StatusUnauthorized, hErr)
	assert.Equal(auth.ErrAPINoSession.Error(), hErr.Message)

	err = s.Login(webtest.Login{
		Username: "kukoon",
		Password: "CHANGEME",
	})
	assert.NoError(err)

	list := []*models.Channel{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/my/channels", nil, http.StatusOK, &list)
	assert.NoError(err)
}
