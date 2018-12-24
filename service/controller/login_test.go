package controller

import (
	"Project/service/model"
	"testing"

	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
)

func Test_Login(t *testing.T) {
	app := newApplogin()
	e := httptest.New(t, app)

	superadmin := model.UserLogin{
		Password: "qwert",
		Username: "super_admin",
	}
	admin := model.UserLogin{
		Password: "qwert",
		Username: "admin",
	}
	admintest := model.LoginTest{
		Username: 2,
	}

	expectedResponse := "json: cannot unmarshal number into Go struct field UserLogin.Username of type string"
	e.POST("/login").WithJSON(superadmin).Expect().Status(httptest.StatusOK).Body().Equal("{\"token\":\"super_admin\"}")
	e.POST("/login").WithJSON(admin).Expect().Status(httptest.StatusOK).Body().Equal("{\"token\":\"admin\"}")
	e.POST("/login").WithJSON(admintest).Expect().Status(httptest.StatusBadRequest).Body().Equal(expectedResponse)
}

func newApplogin() *iris.Application {
	app := iris.New()
	app.Post("/login", Login)
	return app
}
