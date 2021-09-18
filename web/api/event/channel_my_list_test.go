package event

import (
	"net/http"
	"testing"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	"dev.sum7.eu/genofire/golang-lib/web/webtest"
	"github.com/stretchr/testify/assert"

	"github.com/Kukoon/media-server/models"
)

func TestAPIChannelListMy(t *testing.T) {
	assert := assert.New(t)
	s, err := webtest.NewWithDBSetup(bindTest, models.SetupMigration)
	assert.NoError(err)
	defer s.Close()
	assert.NotNil(s)

	hErr := web.HTTPError{}
	// not auth
	err = s.Request(http.MethodGet, "/api/v1/channel/"+models.TestChannelID1.String()+"/events", nil, http.StatusUnauthorized, &hErr)
	assert.NoError(err)
	assert.Equal(auth.ErrAPINoSession.Error(), hErr.Message)

	err = s.Login(webtest.Login{
		Username: "kukoon",
		Password: "CHANGEME",
	})
	assert.NoError(err)

	list := []*models.Event{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/channel/"+models.TestChannelID1.String()+"/events", nil, http.StatusOK, &list)
	assert.NoError(err)
	assert.Len(list, 2)
}
