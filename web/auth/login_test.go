package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Kukoon/media-server/models"
	"github.com/Kukoon/media-server/web"
	"github.com/Kukoon/media-server/web/webtest"
)

func TestAPILogin(t *testing.T) {
	assert := assert.New(t)
	s := webtest.New(assert)
	assert.NotNil(s)

	hErr := web.HTTPError{}
	// invalid
	s.Request(http.MethodPost, "/api/v1/auth/login", 1, http.StatusBadRequest, &hErr)
	assert.Equal(web.APIErrorInvalidRequestFormat, hErr.Message)

	req := login{}
	hErr = web.HTTPError{}
	// invalid - user
	s.Request(http.MethodPost, "/api/v1/auth/login", &req, http.StatusUnauthorized, &hErr)
	assert.Equal(APIErrorUserNotFound, hErr.Message)

	req.Username = "kukoon"
	hErr = web.HTTPError{}
	// invalid - password
	s.Request(http.MethodPost, "/api/v1/auth/login", &req, http.StatusUnauthorized, &hErr)
	assert.Equal(APIErrorIncorrectPassword, hErr.Message)

	req.Password = "CHANGEME"
	obj := models.User{}
	// valid login
	s.Request(http.MethodPost, "/api/v1/auth/login", &req, http.StatusOK, &obj)
	assert.Equal("kukoon", obj.Username)
}
