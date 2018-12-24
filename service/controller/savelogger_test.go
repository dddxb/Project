package controller

import (
	"testing"

	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
)

func Test_SaveLogger(t *testing.T) {
	app := newApplogger()
	e := httptest.New(t, app)

	e.POST("/save_error_logger").WithJSON(iris.Map{"Error": "前端页面展示错误"}).Expect().Status(httptest.StatusOK).Body().Empty()

}

func newApplogger() *iris.Application {
	app := iris.New()
	app.Post("/save_error_logger", SaveLogger)
	return app
}
