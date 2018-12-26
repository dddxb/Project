package controller

import (
	"github.com/kataras/iris"
)

//MessageCount 返回未读消息条数
func MessageCount(ctx iris.Context) {
	ctx.WriteString("3")
}
