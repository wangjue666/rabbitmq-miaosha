package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"rabbitmq-miaosha/web/controllers"
)

func main() {
	//1.创建实例
	app := iris.New()

	//2.设置错误模式，在mvc模式下提示错误
	app.Logger().SetLevel("debug")
	//3.注册模板
	app.RegisterView(iris.HTML("./web/views", ".html"))

	//4.注册控制器
	// 4.5 测试Party的子路由生成规则
	mvc.Configure(app.Party("/hello"), configureMVC)
	//5.启动服务
	app.Run(
		iris.Addr("localhost:8080"),
	)
}

func configureMVC(app *mvc.Application) {
	app.Handle(new(controllers.MovieController))
}
