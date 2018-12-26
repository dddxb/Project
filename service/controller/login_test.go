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

	//合法用户superadmin
	superadmin := model.User{
		Password: "qwert",
		Username: "18753113305",
	}
	//合法用户admin
	admin := model.User{
		Password: "asdfg",
		Username: "18813053143",
	}
	//用户名格式错误
	wrongusernameformat := model.User{
		Password: "asdfg",
		Username: "11011011011",
	}
	//密码格式错误
	wrongpasswordformat := model.User{
		Password: "qwertasdfg", //超出长度限制，设定值为<=6
		Username: "18813053143",
	}
	//密码错误，账号正确
	wrongpassword := model.User{
		Password: "asdfg",
		Username: "18753113305",
	}
	//账号未注册
	wrongusername := model.User{
		Password: "asdfg",
		Username: "18954526289",
	}
	//前端数据未正常读进结构体
	wrongreadjson := model.LoginTest{
		Username: 18753113305,
	}

	e.POST("/login").WithJSON(superadmin).Expect().Status(httptest.StatusOK).Body().Equal("{\"token\":\"super_admin\"}")
	e.POST("/login").WithJSON(admin).Expect().Status(httptest.StatusOK).Body().Equal("{\"token\":\"admin\"}")
	e.POST("/login").WithJSON(wrongusernameformat).Expect().Status(httptest.StatusBadRequest)
	e.POST("/login").WithJSON(wrongpasswordformat).Expect().Status(httptest.StatusBadRequest)
	e.POST("/login").WithJSON(wrongusername).Expect().Status(httptest.StatusBadRequest)
	e.POST("/login").WithJSON(wrongpassword).Expect().Status(httptest.StatusBadRequest)
	e.POST("/login").WithJSON(wrongreadjson).Expect().Status(httptest.StatusBadRequest)
}

func newApplogin() *iris.Application {
	app := iris.New()
	app.Post("/login", Login)
	return app
}
