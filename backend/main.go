package main

import (
	"iris-ticket/backend/models"
	"iris-ticket/backend/routes"

	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	// migrate db
	golog.Info("初始化数据库")
	models.InitDB()
	models.Migration()

	// 加载路由
	routes.RegisterRouter(app)
	err := app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		panic(err)
	}
}
