package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Kukoon/media-server/database"
	"github.com/Kukoon/media-server/models"
	"github.com/Kukoon/media-server/web"
	"github.com/Kukoon/media-server/web/webtest"
)

func TestAPIPassword(t *testing.T) {
	assert := assert.New(t)
	s := webtest.New(assert)
	assert.NotNil(s)
	s.DatabaseMigration(func(db *database.Database) {
		models.SetupMigration(db)
	})

	passwordCurrent := "CHANGEME"
	passwordNew := "test"

	hErr := web.HTTPError{}
	// invalid
	s.Request(http.MethodPost, "/api/v1/my/auth/password", &passwordNew, http.StatusUnauthorized, &hErr)
	assert.Equal(APIErrorNoSession, hErr.Message)

	s.TestLogin()

	res := false
	// set new password
	s.Request(http.MethodPost, "/api/v1/my/auth/password", &passwordNew, http.StatusOK, &res)
	assert.True(res)

	res = false
	// set old password
	s.Request(http.MethodPost, "/api/v1/my/auth/password", &passwordCurrent, http.StatusOK, &res)
	assert.True(res)
}
