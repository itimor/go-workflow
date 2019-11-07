package main

import (
	"iris-ticket/backend/config"
	"iris-ticket/backend/models"
	"iris-ticket/backend/routes"

	"flag"
	"os"

	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	flag.Parse()
	app := newApp()
	err := app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		panic(err)
	}
}

func initDB() {
	models.InitDB()
	models.Migration()
}

func newApp() *iris.Application {
	hostname, _ := os.Hostname()
	if hostname == "wahaha" {
		golog.Info("进入线上环境")
	} else {
		golog.Info("进入测试环境")
	}

	app := iris.New()
	loglevle := config.Conf.Get("test.loglevel").(string)
	app.Logger().SetLevel(loglevle)
	app.Configure(iris.WithOptimizations)

	app.Use(recover.New())
	app.Use(logger.New())

	// app.Logger().SetLevel("debug")
	// app.Use(logger.New())
	// app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
	// 	ctx.JSON(controllers.ApiResource(false, nil, "404 Not Found"))
	// })
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.WriteString("Oh my god, something went wrong, please check code or try again")

	})

	// migrate db
	golog.Info("初始化数据库")
	models.InitDB()
	models.Migration()

	// 加载路由
	routes.RegisterRouter(app)

	//初始化系统 账号 权限 角色
	//models.CreateSystemData(env)

	return app
}
