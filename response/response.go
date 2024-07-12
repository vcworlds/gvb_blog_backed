package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Result(ctx *gin.Context, httpStatus int, code int, msg string, data any) {
	ctx.JSON(httpStatus, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Ok(ctx *gin.Context, msg string, data any) {
	Result(ctx, http.StatusOK, 200, msg, data)
}

func OkWith(ctx *gin.Context) {
	Result(ctx, http.StatusOK, 200, "更改成功", nil)
}

func OkWithMessage(ctx *gin.Context, msg string) {
	Result(ctx, http.StatusOK, 200, msg, map[string]any{})
}

func OkWithData(ctx *gin.Context, data any) {
	Result(ctx, http.StatusOK, 200, "获取成功", data)
}
func Fail(ctx *gin.Context, msg string) {
	Result(ctx, http.StatusUnprocessableEntity, 422, msg, map[string]any{})
}

func FailWithMessage(ctx *gin.Context, msg string) {
	Result(ctx, http.StatusUnprocessableEntity, 422, msg, map[string]any{})
}
