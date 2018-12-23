package controller

import (
	"Project/service/model"
	"fmt"
	"testing"

	"github.com/kataras/iris"
)

func Test_Login(t *testing.T) {
	//app := iris.New()
	e := iris.Handler(Login)

	//e := httptest.New(t, app)
	user := model.UserLogin{
		Password: "qwert",
		Username: "super_admin",
	}

}
