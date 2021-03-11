package webtest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/Kukoon/media-server/models"
	"github.com/Kukoon/media-server/web"
)

type testServer struct {
	gin    *gin.Engine
	ws     *web.Service
	assert *assert.Assertions
	Token  string
}

func New(assert *assert.Assertions) *testServer {
	// db setup
	dbConfig := models.Database{
		Connection: "user=root password=root dbname=media_server host=localhost port=26257 sslmode=disable",
		Testdata:   true,
		Debug:      true,
		LogLevel:   0,
	}
	err := dbConfig.Run()
	assert.Nil(err)
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.NotNil(dbConfig.DB)

	// api setup
	gin.EnableJsonDecoderDisallowUnknownFields()
	gin.SetMode(gin.TestMode)

	ws := &web.Service{
		DB: dbConfig.DB,
	}
	r := gin.Default()
	ws.Bind(r)
	return &testServer{
		gin:    r,
		ws:     ws,
		assert: assert,
	}
}

func (this *testServer) Request(method, url string, body interface{}, expectCode int, jsonObj interface{}) {
	var jsonBody io.Reader
	if body != nil {
		if strBody, ok := body.(string); ok {
			jsonBody = strings.NewReader(strBody)
		} else {
			jsonBodyArray, err := json.Marshal(body)
			this.assert.Nil(err, "no request created")
			jsonBody = bytes.NewBuffer(jsonBodyArray)
		}
	}
	req, err := http.NewRequest(method, url, jsonBody)
	this.assert.Nil(err, "no request created")
	if this.Token != "" {
		req.Header.Set("Authorization", "Bearer "+this.Token)
	}
	w := httptest.NewRecorder()
	this.gin.ServeHTTP(w, req)

	// valid statusCode
	this.assert.Equal(expectCode, w.Code, "expected http status code")
	if expectCode != w.Code {
		fmt.Printf("wrong status code, body:%v\n", w.Body)
		return
	}

	if jsonObj != nil {
		// fetch JSON
		err = json.NewDecoder(w.Body).Decode(jsonObj)
		this.assert.Nil(err, "decode json")
	}
}

/* -
func (this *testServer) login(login login) {
	// POST: correct login
	this.Request(http.MethodPost, "/api/v1/auth/login", &login, http.StatusOK, &this.Token)
	this.assert.NotEqual("", this.Token)
}
*/
