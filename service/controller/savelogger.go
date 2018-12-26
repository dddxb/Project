package controller

import (
	"Project/service/model"
	"fmt"

	"github.com/kataras/iris"
)

//SaveLogger 存储日志
func SaveLogger(ctx iris.Context) {
	var err model.Errlogger

	ctx.ReadJSON(&err)
	if err.Error != "" {
		ctx.WriteString("success")
	}
	fmt.Println(err)
}
