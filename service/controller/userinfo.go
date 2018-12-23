package controller

import (
	"Project/service/model"
	"fmt"
	"github.com/kataras/iris"
)


func UserInfo(ctx iris.Context) {
	fmt.Println(ctx.Path())
	//request url:localhost:8080/get_info?token=super_admin或者admin
	user := ctx.URLParam("token")
	fmt.Println(user)

	super_admin := model.UserInfo{
		Name:      "super_admin",
		User_id:    1,
		Access:    []string{"super_admin","admin"},
		Token:     "super_admin",
		Avator:    "https://file.iviewui.com/dist/a0e88e83800f138b94d2414621bd9704.png",
	}
	admin := model.UserInfo{
		Name:      "admin",
		User_id:    2,
		Access:    []string{"admin"},
		Token:     "admin",
		Avator:    "https://avatars0.githubusercontent.com/u/20942571?s=460&v=4",
	}

	switch user {
	case "super_admin":
		ctx.JSON(super_admin)
	case "admin":
		ctx.JSON(admin)
	}
}


























//var usr model.UserInfo

//super_admin := user{"super_admin",1,[]string{"super_admin","admin"},"super_admin","https://file.iviewui.com/dist/a0e88e83800f138b94d2414621bd9704.png"}
//if usr.Username == "super_admin" {
//	//结构体转Json
//	back, _ := ctx.JSON(super_admin)
//
//	ctx.Writef("Received: %#+v\n", back)
//}