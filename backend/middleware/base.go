package middleware

import (
	"fmt"
	"iris-ticket/backend/controllers/common"
	"strings"

	"github.com/kataras/iris"
)

// NoMethodHandler 未找到请求方法的处理函数
func NoMethodHandler() iris.Handler {
	return func(ctx iris.Context) {
		common.ResFail(ctx, "方法不被允许")
	}
}

// NoRouteHandler 未找到请求路由的处理函数
func NoRouteHandler() iris.Handler {
	return func(ctx iris.Context) {
		common.ResFail(ctx, "未找到请求路由的处理函数")
	}
}

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(iris.Context) bool

// AllowPathPrefixSkipper 检查请求路径是否包含指定的前缀，如果包含则跳过
func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(ctx iris.Context) bool {
		path := ctx.Path()
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// AllowPathPrefixNoSkipper 检查请求路径是否包含指定的前缀，如果包含则不跳过
func AllowPathPrefixNoSkipper(prefixes ...string) SkipperFunc {
	return func(ctx iris.Context) bool {
		path := ctx.Path()
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return false
			}
		}
		return true
	}
}

// AllowMethodAndPathPrefixSkipper 检查请求方法和路径是否包含指定的前缀，如果包含则跳过
func AllowMethodAndPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(ctx iris.Context) bool {
		path := JoinRouter(ctx.Method(), ctx.Path())
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// JoinRouter 拼接路由
func JoinRouter(method, path string) string {
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", strings.ToUpper(method), path)
}
