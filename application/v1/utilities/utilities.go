package OTUtilities

import (
	"encoding/json"
	"strconv"

	"github.com/valyala/fasthttp"
)

func GetJsonHttpHeader(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "application/json; charset=utf-8")
	ctx.Response.Header.Set("X-Frame-Options", "SAMEORIGIN")
	ctx.Response.Header.Set("X-Content-Type-Options", "nosniff")
}

func GetJsonStringFromStruct(class interface{}) string {
	jsonByte, _ := json.Marshal(class)

	return string(jsonByte)
}

func GetUserValueAsString(value string, ctx *fasthttp.RequestCtx) string {
	return ctx.UserValue(value).(string)
}

func GetUserValueAsInt(value string, ctx *fasthttp.RequestCtx) int {
	returnValue, _ := strconv.Atoi(ctx.UserValue(value).(string))
	return returnValue
}
