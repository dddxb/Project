package controller

import (
	"fmt"

	"github.com/kataras/iris"
)

//Logout 退出登录，无返回值
func Logout(ctx iris.Context) {

	fmt.Println(ctx)

}
