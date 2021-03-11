package docs

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Kukoon/media-server/web/webtest"
)

func TestMetricsLoaded(t *testing.T) {
	assert := assert.New(t)
	s := webtest.New(assert)
	assert.NotNil(s)

	// GET
	s.Request(http.MethodGet, "/metrics", nil, http.StatusOK, nil)
}
