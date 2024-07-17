package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/global"
	"gvb_blog/models"
	"gvb_blog/models/ctype"
	"gvb_blog/response"
	"gvb_blog/service/user_service"
	"gvb_blog/utils"
	"gvb_blog/utils/jwt"
)

func (UserApi) Login(ctx *gin.Context) {
	var userRes user_service.UserResponse
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
	if !utils.ValidPassword(userRes.Password, "PasswordSalt", userMo.Password) {
		response.Fail(ctx, "您的密码输入有误")
		return
	}
	// 携带token
	token, err := jwt.ReleaseToken(userMo)
	if err != nil {
		global.Log.Error(err)
		response.Fail(ctx, "未知的错误")
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
	if err == nil {
		response.Fail(ctx, "该用户已存在")
		return
	}
	// 判断两次密码是否一致
	if registerRes.RePassword != registerRes.Password {
		response.Fail(ctx, "两次密码不一致")
		return
	}
	hashPassword := utils.EncryptPassword(registerRes.Password, "PasswordSalt")
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
