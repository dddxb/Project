package controller

import (
	"testing"

	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
)

func Test_MessageCount(t *testing.T) {
	app := newAppcount()
	e := httptest.New(t, app)

	e.GET("/message/count").Expect().Status(httptest.StatusOK).Body().Equal("3")

}

func newAppcount() *iris.Application {
	app := iris.New()
	app.Get("/message/count", MessageCount)
	return app
}
