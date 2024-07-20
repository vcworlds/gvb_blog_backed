package user_api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gvb_blog/common"
	"gvb_blog/dao"
	"gvb_blog/global"
	"gvb_blog/models"
	"gvb_blog/response"
	"gvb_blog/service/user_service"
	"gvb_blog/utils/jwt"
)

func (UserApi) Login(ctx *gin.Context) {
	var userRes user_service.LoginService
	err := ctx.ShouldBindJSON(&userRes)
	if err != nil {
		response.FailWithValidateError(err, &userRes, ctx)
		return
	}
	res := userRes.LoginService()
	if res.Code != 200 {
		response.Fail(ctx, res.Msg)
		return
	}
	response.Ok(ctx, "登录成功", gin.H{"token": res.Data})

}

func (UserApi) Register(ctx *gin.Context) {
	var registerRes user_service.RegisterService
	err := ctx.ShouldBindJSON(&registerRes)
	if err != nil {
		response.FailWithValidateError(err, &registerRes, ctx)
		return
	}
	res := registerRes.RegisterService()
	if res.Code != 200 {
		response.Fail(ctx, res.Msg)
	}
	response.OkWithMessage(ctx, res.Msg)
}

func (UserApi) UserList(ctx *gin.Context) {
	var page models.Page
	err := ctx.ShouldBindQuery(&page)
	if err != nil {
		response.Fail(ctx, "分页条件绑定失败")
		return
	}
	_claims, exits := ctx.Get("claims")
	if !exits {
		response.Fail(ctx, "token不存在")
		return
	}
	res := user_service.UserListService(page, _claims)
	if res.Code != 200 {
		response.Fail(ctx, res.Msg)
	}
	// 分页获取
	response.OkWithData(ctx, res.Data)

}

func (UserApi) RoleUpdate(ctx *gin.Context) {
	var RoleRes user_service.RoleUpdateService
	err := ctx.ShouldBindJSON(&RoleRes)
	if err != nil {
		global.Log.Error(err)
		response.FailWithValidateError(err, &RoleRes, ctx)
		return
	}
	var userMo models.UserModel
	err = global.DB.Take(&userMo, RoleRes.UserId).Error
	if err != nil {
		global.Log.Error(err)
		response.Fail(ctx, "查询用户失败")
		return
	}
	err = global.DB.Model(&userMo).Update("role", RoleRes.Role).Error
	if err != nil {
		global.Log.Error(err)
		response.Fail(ctx, "更新用户权限失败")
		return
	}
	response.OkWithMessage(ctx, "更新成功")
}

// PasswordUpdate 用户修改密码
func (UserApi) PasswordUpdate(ctx *gin.Context) {
	var PasswordRep user_service.PasswordUpdateService
	err := ctx.ShouldBindJSON(&PasswordRep)
	if err != nil {
		global.Log.Error(err)
		response.FailWithValidateError(err, &PasswordRep, ctx)
		return
	}
	_claims, exit := ctx.Get("claims")
	if !exit {
		response.Fail(ctx, "token有误")
		return
	}
	res := PasswordRep.RegisterService(_claims)
	if res.Code != 200 {
		response.Fail(ctx, res.Msg)
		return
	}
	response.OkWithMessage(ctx, res.Msg)
}

// 注销登录，将token变为无效
func (receiver UserApi) UserLogout(ctx *gin.Context) {
	_claims, exits := ctx.Get("claims")
	if !exits {
		response.Fail(ctx, "token不存在")
		return
	}
	claims := _claims.(*jwt.Claims)
	tokenKey := fmt.Sprintf("token_%d", claims.UserId)
	_, err := global.Redis.Get(context.Background(), tokenKey).Result()
	if err == redis.Nil {
		global.Log.Error(err)
		response.Fail(ctx, "token不存在")
		ctx.Abort()
		return
	} else if err != nil {
		global.Log.Error(err)
		response.Fail(ctx, "token不存在")
		ctx.Abort()
		return
	}
	// 将 token 过期时间设置为0，几乎立即失效
	err = global.Redis.Expire(context.Background(), tokenKey, 0).Err()
	if err != nil {
		global.Log.Error("Error setting token expiration:", err)
		response.Fail(ctx, "注销失败")
		return
	}
	response.OkWithMessage(ctx, "注销成功")
}

// 删除用户
func (UserApi) UserDelete(ctx *gin.Context) {
	var ids common.RemoveFileList
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithValidateError(err, &ids, ctx)
		return
	}
	// 判断id是否存在
	var userMo []models.UserModel
	err = global.DB.Find(&userMo, ids.Ids).Error
	if err != nil {
		global.Log.Error(err)
		response.Fail(ctx, fmt.Sprintf("用户id：%d不存在", ids.Ids))
		return
	}
	count, err := dao.DeleteCommon[models.UserModel](userMo, ids.Ids)
	if err != nil {
		global.Log.Error(err)
		response.Fail(ctx, "删除失败")
		return
	}
	response.OkWithMessage(ctx, fmt.Sprintf("删除成功，共删除%d个用户", count))
}
