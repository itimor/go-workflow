package routes

import (
	"iris-ticket/backend/config"
	"iris-ticket/backend/middleware"
	"os"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

// 所有的路由
func RegisterRouter(app *iris.Application) {
	preSettring(app)
	main := corsSetting(app)

	// 加载路由
	AuthRoute(main) // 认证登录
	SysRoute(main) // 系统管理

}

func corsSetting(app *iris.Application) (main iris.Party) {

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})

	/* 定义路由 */
	main = app.Party("/", crs).AllowMethods(iris.MethodOptions)
	main.Use(middleware.ServeHTTP)

	main.Get("/", func(ctx iris.Context) { // 首页模块
		//_ = ctx.View("index.html")
		ctx.HTML("<h1 style='height: 1000px;line-height: 1000px;text-align: center;'>召唤师，欢迎来到王者峡谷</h1>")
	})

	return main
}

func preSettring(app *iris.Application) {
	hostname, _ := os.Hostname()
	if hostname == "wahaha" {
		golog.Info("进入线上环境")
	} else {
		golog.Info("进入测试环境")
	}

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

	// ---------------------- 自定义错误处理 ------------------------
	// app.OnErrorCode(iris.StatusNotFound, logger, func(ctx iris.Context) {
	// 	supports.Error(ctx, iris.StatusNotFound, supports.NotFound, nil)
	// })
	// app.OnErrorCode(iris.StatusInternalServerError, logger, func(ctx iris.Context) {
	// 	supports.Error(ctx, iris.StatusInternalServerError, supports.StatusInternalServerError, nil)
	// })
}
