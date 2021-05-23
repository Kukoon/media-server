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

	"github.com/Kukoon/media-server/database"
	"github.com/Kukoon/media-server/web"
)

type testServer struct {
	db          *database.Database
	gin         *gin.Engine
	ws          *web.Service
	assert      *assert.Assertions
	lastCookies []*http.Cookie
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func New(assert *assert.Assertions) *testServer {
	// db setup
	dbConfig := database.Database{
		Connection: "user=root password=root dbname=media_server host=localhost port=26257 sslmode=disable",
		Testdata:   true,
		Debug:      false,
		LogLevel:   0,
	}
	err := dbConfig.Run()
	if err != nil && err != database.ErrNothingToMigrate {
		fmt.Println(err.Error())
		assert.Nil(err)
	}
	assert.NotNil(dbConfig.DB)

	// api setup
	gin.EnableJsonDecoderDisallowUnknownFields()
	gin.SetMode(gin.TestMode)

	ws := &web.Service{
		DB: dbConfig.DB,
	}
	ws.Session.Name = "mysession"
	ws.Session.Secret = "hidden"

	r := gin.Default()
	ws.LoadSession(r)
	ws.Bind(r)
	return &testServer{
		db:     &dbConfig,
		gin:    r,
		ws:     ws,
		assert: assert,
	}
}
func (this *testServer) DatabaseMigration(f func(db *database.Database)) {
	f(this.db)
	this.db.MigrateTestdata()
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
	if len(this.lastCookies) > 0 {
		for _, c := range this.lastCookies {
			req.AddCookie(c)
		}
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

	result := w.Result()
	if result != nil {
		cookies := result.Cookies()
		if len(cookies) > 0 {
			this.lastCookies = cookies
		}
	}
}

func (this *testServer) Login(login Login) {
	// POST: correct login
	this.Request(http.MethodPost, "/api/v1/auth/login", &login, http.StatusOK, nil)
}

func (this *testServer) TestLogin() {
	this.Login(Login{
		Username: "kukoon",
		Password: "CHANGEME",
	})
}
