package main

import (
	"Project/service/model"
	"testing"

	"github.com/kataras/iris/httptest"
)

func Test_newApp(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)

	user := model.UserLogin{
		Password: "qwert",
		Username: "super_admin",
	}

	e.POST("/login").WithJSON(user).Expect().Status(httptest.StatusOK).Body().Equal("{\"token\":\"super_admin\"}")
	e.POST("/save_error_logger").Expect().Status(httptest.StatusOK).Body().Empty()
	e.GET("/message/count").Expect().Status(httptest.StatusOK).Body().Equal("3")
	expectedResponse := "{\"username\":\"super_admin\",\"user_id\":1,\"access\":[\"super_admin\",\"admin\"],\"token\":\"super_admin\",\"avator\":\"https://file.iviewui.com/dist/a0e88e83800f138b94d2414621bd9704.png\"}"
	e.GET("/get_info").WithURL("localhost:8080/get_info?token=super_admin").Expect().Status(httptest.StatusOK).Body().Equal(expectedResponse)

}
