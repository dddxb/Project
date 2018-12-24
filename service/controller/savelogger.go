package controller

import (
	"Project/service/model"
	"fmt"

	"github.com/kataras/iris"
)

func SaveLogger(ctx iris.Context) {
	var err model.Errlogger

	ctx.ReadJSON(&err)

	fmt.Println(err)
}
