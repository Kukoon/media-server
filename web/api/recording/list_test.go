package recording

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Kukoon/media-server/models"
	// "github.com/Kukoon/media-server/web"
	"github.com/Kukoon/media-server/web/webtest"
)

func TestAPIRecordingsList(t *testing.T) {
	assert := assert.New(t)
	s := webtest.New(assert)
	assert.NotNil(s)

	list := []*models.Recording{}
	// GET
	s.Request(http.MethodGet, "/api/v1/recordings", nil, http.StatusOK, &list)
}
