package controller

import (
	"github.com/kataras/iris"
)

func MessageCount(ctx iris.Context) {

	ctx.WriteString("3")

}