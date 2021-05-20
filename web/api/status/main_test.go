package status

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Kukoon/media-server/web/webtest"
)

func TestAPIStatus(t *testing.T) {
	assert := assert.New(t)
	s := webtest.New(assert)
	assert.NotNil(s)

	obj := Status{}
	// GET - common name
	s.Request(http.MethodGet, "/api/status", nil, http.StatusOK, &obj)
	assert.Equal(VERSION, obj.Version)
	assert.Equal(EXTRAS, obj.Extras)
	assert.True(obj.Up)
}
