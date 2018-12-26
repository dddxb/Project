package controller

import (
	"testing"

	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
)

func Test_UserInfo(t *testing.T) {
	app := newAppinfo()
	e := httptest.New(t, app)

	expectedResponseadmin := "{\"password\":\"\",\"username\":\"18813053143\",\"user_id\":\"2\",\"access\":\"admin\",\"token\":\"admin\",\"avator\":\"https://avatars0.githubusercontent.com/u/20942571?s=460\\u0026v=4\"}"
	e.GET("/get_info").WithURL("localhost:8080/get_info?token=admin").Expect().Status(httptest.StatusOK).Body().Equal(expectedResponseadmin)
	expectedResponsesuper := "{\"password\":\"\",\"username\":\"18753113305\",\"user_id\":\"1\",\"access\":\"super_admin\",\"token\":\"super_admin\",\"avator\":\"https://file.iviewui.com/dist/a0e88e83800f138b94d2414621bd9704.png\"}"
	e.GET("/get_info").WithURL("localhost:8080/get_info?token=super_admin").Expect().Status(httptest.StatusOK).Body().Equal(expectedResponsesuper)
}

func newAppinfo() *iris.Application {
	app := iris.New()
	app.Get("/get_info", UserInfo)
	return app
}
