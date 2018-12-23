package controller

import (
	"Project/service/model"
	"fmt"
	"github.com/kataras/iris"
)

func Login(ctx iris.Context) {
	var usr model.UserLogin

	if err := ctx.ReadJSON(&usr); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	fmt.Println(usr.Username,usr.Password)

	switch usr.Username {
	case "super_admin" :
		ctx.JSON(iris.Map{"token": "super_admin"})
	case "admin":
		ctx.JSON(iris.Map{"token":"admin"})
	}

}



































//super_admin := user{"super_admin",1,[]string{"super_admin","admin"},"super_admin","https://file.iviewui.com/dist/a0e88e83800f138b94d2414621bd9704.png"}
//if usr.Username == "super_admin" {
//	//结构体转Json
//	back, _ := ctx.JSON(super_admin)
//
//
//	ctx.Writef("Received: %#+v\n", back)
//}