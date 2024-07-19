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
	"gvb_blog/models/ctype"
	"gvb_blog/response"
	"gvb_blog/service/user_service"
	"gvb_blog/utils"
	"gvb_blog/utils/jwt"
	"log"
	"time"
)

func (UserApi) Login(ctx *gin.Context) {
	var userRes user_service.LoginResponse
	err := ctx.ShouldBindJSON(&userRes)
	if err != nil {
		response.FailWithValidateError(err, &userRes, ctx)
		return
	}
	// 判断用户名是否存在
	var userMo models.UserModel
	err = global.DB.Take(&userMo, "user_name = ?", userRes.UserName).Error
	if err != nil {
		response.Fail(ctx, "未找到该用户")
		return
	}
	// 判断密码是否正确
	exits, err := dao.IsPassword(userRes.Password, userMo.ID)
	if !exits || err != nil {
		global.Log.Error(err)
		response.Fail(ctx, "您的密码输入有误")
		return
	}
	// 携带token
	token, err := jwt.ReleaseToken(userMo)
	if err != nil {
		global.Log.Error(err)
		response.Fail(ctx, "未知的错误,请刷新重试")
		return
	}
	// 将token存储到redis
	expiration := global.Config.Jwt.Expires * 60 * 60
	id := fmt.Sprintf("token_%d", userMo.ID)

	if global.Redis == nil {
		log.Println("global.Redis is nil")
		response.Fail(ctx, "Redis 初始化失败")
		return
	}

	err = global.Redis.Set(context.Background(), id, token, time.Duration(expiration)*time.Second).Err()
	if err != nil {
		// 处理错误
		log.Println("Error setting token in Redis:", err)
		response.Fail(ctx, "存储 Token 失败")
		return
	}
	response.Ok(ctx, "登录成功", gin.H{"token": token})

}

func (UserApi) Register(ctx *gin.Context) {
	var registerRes user_service.RegisterRep
	err := ctx.ShouldBindJSON(&registerRes)
	if err != nil {
		response.FailWithValidateError(err, &registerRes, ctx)
		return
	}
	// 判断用户名是否已经注册
	var userMo models.UserModel
	err = global.DB.Take(&userMo, "user_name = ?", registerRes.UserName).Error
	fmt.Println(registerRes.Salt)
	if err == nil {
		response.Fail(ctx, "该用户已存在")
		return
	}
	// 判断两次密码是否一致
	if registerRes.RePassword != registerRes.Password {
		response.Fail(ctx, "两次密码不一致")
		return
	}
	// 判断邮箱
	exits := utils.IsValidEmail(registerRes.Email)
	if !exits {
		response.Fail(ctx, "邮箱不合规范")
		return
	}
	if registerRes.Salt == "" {
		registerRes.Salt = "PasswordSalt"
	}
	hashPassword := utils.EncryptPassword(registerRes.Password, registerRes.Salt)
	global.DB.Create(&models.UserModel{
		NickName:   registerRes.NickName,
		UserName:   registerRes.UserName,
		Password:   hashPassword,
		AvatarId:   1,
		Email:      registerRes.Email,
		Addr:       "暂未开发",
		IP:         "暂未开放",
		Role:       ctype.PermissionUser,
		SignStatus: ctype.SignEmail,
	})
	response.OkWithMessage(ctx, "注册成功")
}

func (UserApi) UserList(ctx *gin.Context) {
	var page models.Page
	err := ctx.ShouldBindQuery(&page)
	if err != nil {
		response.Fail(ctx, "分页条件绑定失败")
		return
	}
	option := common.Option{page}
	userList, _, err := common.CommonPage[models.UserModel](models.UserModel{}, option)
	if err != nil {
		global.Log.Error(err)
		response.Fail(ctx, "获取列表失败")
	}
	_claims, exits := ctx.Get("claims")
	if !exits {
		response.Fail(ctx, "token不存在")
		return
	}
	claims := _claims.(*jwt.Claims)
	var userModel []models.UserModel
	for _, user := range userList {
		if claims.Role != "管理员" {
			user.UserName = ""
			user.Salt = "*****"
			user.Tel = utils.DesensitizationPhone(user.Tel)
			user.Email = utils.DesensitizationEmail(user.Email)
		}
		userModel = append(userModel, user)
	}
	// 分页获取
	response.OkWithData(ctx, userModel)

}

func (UserApi) RoleUpdate(ctx *gin.Context) {
	var RoleRes user_service.RoleUpdateRep
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
	var PasswordRep user_service.PasswordUpdateRep
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
	claims := _claims.(*jwt.Claims)
	var userMo models.UserModel
	err = global.DB.Take(&userMo, claims.UserId).Error
	if err != nil {
		response.Fail(ctx, "系统错误，请刷新重试")
		return
	}
	// 判断密码是否正确
	exits, err := dao.IsPassword(PasswordRep.OldPassword, userMo.ID)
	if err != nil || !exits {
		response.Fail(ctx, "原始密码错误")
		return
	}
	// 修改密码
	newPassword := utils.EncryptPassword(PasswordRep.NewPassword, userMo.Salt)
	err = global.DB.Model(&userMo).Update("password", newPassword).Error
	if err != nil {
		response.Fail(ctx, "修改失败")
		return
	}
	response.OkWithMessage(ctx, "修改成功")
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
