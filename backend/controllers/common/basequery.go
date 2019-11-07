package common

import (
	"errors"

	"iris-ticket/backend/pkg/convert"

	"github.com/kataras/iris"
)

// GetQueryToStrE
func GetQueryToStrE(ctx iris.Context, key string) (string, error) {
	ok := ctx.URLParamExists(key)
	if !ok {
		return "", errors.New("没有这个值传入")
	}
	return ctx.URLParam(key), nil
}

// GetQueryToStr
func GetQueryToStr(ctx iris.Context, key string, defaultValues ...string) string {
	var defaultValue string
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	str, err := GetQueryToStrE(ctx, key)

	if str == "" || err != nil {
		return defaultValue
	}
	return str
}

// QueryToUintE
func GetQueryToUintE(ctx iris.Context, key string) (uint, error) {
	str, err := GetQueryToStrE(ctx, key)
	if err != nil {
		return 0, err
	}
	return convert.ToUintE(str)
}

// QueryToUint
func GetQueryToUint(ctx iris.Context, key string, defaultValues ...uint) uint {
	var defaultValue uint
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	val, err := GetQueryToUintE(ctx, key)
	if err != nil {
		return defaultValue
	}
	return val
}

// QueryToUintE
func GetQueryToUint64E(ctx iris.Context, key string) (uint64, error) {
	str, err := GetQueryToStrE(ctx, key)
	if err != nil {
		return 0, err
	}
	return convert.ToUint64E(str)
}

// QueryToUint
func GetQueryToUint64(ctx iris.Context, key string, defaultValues ...uint64) uint64 {
	var defaultValue uint64
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	val, err := GetQueryToUint64E(ctx, key)
	if err != nil {
		return defaultValue
	}
	return val
}
