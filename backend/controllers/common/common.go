package common

import (
	"github.com/kataras/iris"
)

const (
	SUCCESS_CODE          = 20000      //成功的状态码
	FAIL_CODE             = 30000      //失败的状态码
	MD5_PREFIX            = "jkfldfsf" //MD5加密前缀字符串
	TOKEN_KEY             = "X-Token"  //页面token键名
	USER_ID_Key           = "X-USERID" //页面用户ID键名
	USER_UUID_Key         = "X-UUID"   //页面UUID键名
	SUPER_ADMIN_ID uint64 = 1          // 超级管理员账号ID
)

type ResponseModel struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponseModelBase struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 响应JSON数据
func ResJSON(ctx iris.Context, status int, v interface{}) {
	ctx.StatusCode(status)
	ctx.JSON(v)
}

// 响应成功
func ResSuccess(ctx iris.Context, v interface{}) {
	ret := ResponseModel{Code: SUCCESS_CODE, Msg: "ok", Data: v}
	ResJSON(ctx, iris.StatusOK, &ret)
}

// 响应成功
func ResSuccessMsg(ctx iris.Context) {
	ret := ResponseModelBase{Code: SUCCESS_CODE, Msg: "ok"}
	ResJSON(ctx, iris.StatusOK, &ret)
}

// 响应失败
func ResFail(ctx iris.Context, msg string) {
	ret := ResponseModelBase{Code: FAIL_CODE, Msg: msg}
	ResJSON(ctx, iris.StatusOK, &ret)
}

// 响应失败
func ResFailCode(ctx iris.Context, msg string, code int) {
	ret := ResponseModelBase{Code: code, Msg: msg}
	ResJSON(ctx, iris.StatusOK, &ret)
}

// 响应错误-服务端故障
func ResErrSrv(ctx iris.Context, err error) {
	ret := ResponseModelBase{Code: FAIL_CODE, Msg: "服务端故障"}
	ResJSON(ctx, iris.StatusOK, &ret)
}

// 响应错误-用户端故障
func ResErrCli(ctx iris.Context, err error) {
	ret := ResponseModelBase{Code: FAIL_CODE, Msg: "err"}
	ResJSON(ctx, iris.StatusOK, &ret)
}

type ResponsePageData struct {
	Total uint64      `json:"total"`
	Items interface{} `json:"items"`
}

type ResponsePage struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data ResponsePageData `json:"data"`
}

// 响应成功-分页数据
func ResSuccessPage(ctx iris.Context, total uint64, list interface{}) {
	ret := ResponsePage{Code: SUCCESS_CODE, Msg: "ok", Data: ResponsePageData{Total: total, Items: list}}
	ResJSON(ctx, iris.StatusOK, &ret)
}

// 获取页码
func GetPageIndex(ctx iris.Context) uint64 {
	return GetQueryToUint64(ctx, "page", 1)
}

// 获取每页记录数
func GetPageLimit(ctx iris.Context) uint64 {
	limit := GetQueryToUint64(ctx, "limit", 10)
	if limit > 100 {
		limit = 10
	}
	return limit
}

// 获取排序信息
func GetPageSort(ctx iris.Context) string {
	return GetQueryToStr(ctx, "sort")
}

// 获取搜索关键词信息
func GetPageKey(ctx iris.Context) string {
	return GetQueryToStr(ctx, "key")
}
