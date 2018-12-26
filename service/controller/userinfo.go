package controller

import (
	"Project/service/biz"
	"fmt"

	"github.com/kataras/iris"
)

// UserInfo UserInfo
func UserInfo(ctx iris.Context) {
	//request url:localhost:8080/get_info?token=super_admin或者admin
	user := ctx.URLParam("token")
	fmt.Println(user)

	//查库where TOKEN=token
	res := biz.UserInfo(user)
	ctx.JSON(res)
}
