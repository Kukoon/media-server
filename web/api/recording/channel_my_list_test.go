package recording

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
	err = s.Request(http.MethodGet, "/api/v1/channel/"+models.TestChannelID1.String()+"/recordings", nil, http.StatusUnauthorized, &hErr)
	assert.NoError(err)
	assert.Equal(auth.ErrAPINoSession.Error(), hErr.Message)

	err = s.Login(webtest.Login{
		Username: "kukoon",
		Password: "CHANGEME",
	})
	assert.NoError(err)

	list := []*models.Recording{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/channel/"+models.TestChannelID1.String()+"/recordings", nil, http.StatusOK, &list)
	assert.NoError(err)
	assert.Len(list, 22)
	if len(list) > 0 {
		obj := list[0]
		assert.Nil(obj.Lang)
	}

	list = []*models.Recording{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/channel/"+models.TestChannelID1.String()+"/recordings?lang=de", nil, http.StatusOK, &list)
	assert.NoError(err)
	assert.Len(list, 22)
	if len(list) > 0 {
		obj := list[0]
		assert.NotNil(obj.Lang)
	}

	hErr = web.HTTPError{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/channel/"+models.TestChannelID1.String()+"/recordings?event=a", nil, http.StatusBadRequest, &hErr)
	assert.NoError(err)
	assert.Equal(web.ErrAPIInvalidRequestFormat.Error(), hErr.Message)

	list = []*models.Recording{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/channel/"+models.TestChannelID1.String()+"/recordings?event="+models.TestEventID2.String(), nil, http.StatusOK, &list)
	assert.NoError(err)
	assert.Len(list, 8)

	hErr = web.HTTPError{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/channel/"+models.TestChannelID1.String()+"/recordings?tag=a", nil, http.StatusBadRequest, &hErr)
	assert.NoError(err)
	assert.Equal(web.ErrAPIInvalidRequestFormat.Error(), hErr.Message)

	list = []*models.Recording{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/channel/"+models.TestChannelID1.String()+"/recordings?tag="+models.TestTagBuchvorstellungID.String(), nil, http.StatusOK, &list)
	assert.NoError(err)
	assert.Len(list, 6)

	list = []*models.Recording{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/channel/"+models.TestChannelID1.String()+"/recordings?tag="+models.TestTagBuchvorstellungID.String()+"&tag="+models.TestTagDiskussionID.String(), nil, http.StatusOK, &list)
	assert.NoError(err)
	assert.Len(list, 5)

	hErr = web.HTTPError{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/channel/"+models.TestChannelID1.String()+"/recordings?speaker=a", nil, http.StatusBadRequest, &hErr)
	assert.NoError(err)
	assert.Equal(web.ErrAPIInvalidRequestFormat.Error(), hErr.Message)

	list = []*models.Recording{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/channel/"+models.TestChannelID1.String()+"/recordings?speaker="+models.TestStream1IDSpeaker1.String(), nil, http.StatusOK, &list)
	assert.NoError(err)
	assert.Len(list, 1)

	hErr = web.HTTPError{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/channel/"+models.TestChannelID1.String()+"/recordings?channel=a", nil, http.StatusBadRequest, &hErr)
	assert.NoError(err)
	assert.Equal(web.ErrAPIInvalidRequestFormat.Error(), hErr.Message)

	list = []*models.Recording{}
	// GET
	err = s.Request(http.MethodGet, "/api/v1/channel/"+models.TestChannelID1.String()+"/recordings?channel="+models.TestChannelID1.String(), nil, http.StatusOK, &list)
	assert.NoError(err)
	assert.Len(list, 22)
}
