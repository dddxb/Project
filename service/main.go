package main

import (
	"Project/service/controller"
	"github.com/kataras/iris/middleware/logger"

	"github.com/kataras/iris"
)

func main() {
	app := newApp()

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}

func newApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(logger.New())

	//访问静态资源
	app.StaticWeb("/", "./webui/dist")

	//POST {username，password}
	// response token
	app.Post("/login", controller.Login)

	//request url:localhost:8080/get_info?token=super_admin或者admin
	//response userinfo struct
	app.Get("/get_info", controller.UserInfo)

	//退出登录，无response
	app.Post("/logout", controller.Logout)

	//后台存储日志
	app.Post("/save_error_logger", controller.SaveLogger)

	//未读消息个数，response 3
	app.Get("/message/count", controller.MessageCount)

	return app
}
