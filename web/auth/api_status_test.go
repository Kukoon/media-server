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

func TestAPIStatus(t *testing.T) {
	assert := assert.New(t)
	s := webtest.New(assert)
	assert.NotNil(s)
	s.DatabaseMigration(func(db *database.Database) {
		models.SetupMigration(db)
	})

	hErr := web.HTTPError{}
	// invalid
	s.Request(http.MethodGet, "/api/v1/auth/status", nil, http.StatusUnauthorized, &hErr)
	assert.Equal(APIErrorNoSession, hErr.Message)

	s.TestLogin()

	obj := models.User{}
	// invalid - user
	s.Request(http.MethodGet, "/api/v1/auth/status", nil, http.StatusOK, &obj)
	assert.Equal("kukoon", obj.Username)

}
