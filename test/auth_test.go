package test

import (
	"gozzafadillah/config"
	"gozzafadillah/controllers"
	"gozzafadillah/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Cara test
// go test ./... -v -coverpkg=./controllers/...,./lib/...,./models/... -coverprofile=cover.out && go tool cover -html=cover.out

type TestCase struct {
	Method         string
	Name           string
	Path           string
	ExpectStatus   int
	ExpectResponse string
}

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	config.InitMigrateTest()
	e := echo.New()
	return e
}
func InsertDataUser() error {
	user := models.User{
		ID:       1,
		Username: "gozzafadillah",
		Password: "12345",
	}

	var err error
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
func TestRegister(t *testing.T) {
	SuccessTest := TestCase{

		Method:         http.MethodPost,
		Name:           "Register User",
		Path:           "/register",
		ExpectStatus:   http.StatusOK,
		ExpectResponse: "Register Success !",
	}

	e := InitEchoTestAPI()
	reqStr := `{
		"ID":2,
		"username": "KianaSekar",
		"password": "12345"
	}`

	req := httptest.NewRequest(SuccessTest.Method, SuccessTest.Path, strings.NewReader(reqStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(SuccessTest.Path)
	if assert.NoError(t, controllers.RegisterController(c)) {
		assert.Equal(t, SuccessTest.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), SuccessTest.ExpectResponse)

	}
}
func TestLogin(t *testing.T) {
	successTest := TestCase{

		Method:         http.MethodPost,
		Name:           "Login User",
		Path:           "/login",
		ExpectStatus:   http.StatusOK,
		ExpectResponse: "Login Success !",
	}

	e := InitEchoTestAPI()
	InsertDataUser()
	reqStr := `{
		"Username": "gozzafadillah",
		"password": "12345"
	}`

	req := httptest.NewRequest(successTest.Method, successTest.Path, strings.NewReader(reqStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(successTest.Path)
	if assert.NoError(t, controllers.LoginController(c)) {
		assert.Equal(t, successTest.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), successTest.ExpectResponse)

	}
}
