package response

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/utils"
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
func OkWithPage(ctx *gin.Context, list any, count int64) {
	Result(ctx, http.StatusOK, 200, "获取成功", gin.H{"count": count, "list": list})
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

func FailWithValidateError[T any](err error, obj *T, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	Fail(c, msg)
}
