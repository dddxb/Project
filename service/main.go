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

	app.StaticWeb("/", "./webui/dist")

	app.Get("/", func(ctx iris.Context) {
		ctx.ServeFile("./webui/dist/index.html", false)
	})

	//
	app.Post("/login", controller.Login)
	app.Get("/get_info", controller.UserInfo)
	app.Post("/logout", controller.Logout)



	app.Post("/save_error_logger", controller.SaveLogger)

	app.Get("/message/count", controller.MessageCount)

	return app
}

