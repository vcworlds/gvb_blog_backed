package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_blog/common"
	"gvb_blog/dao"
	"gvb_blog/global"
	"gvb_blog/models"
	"gvb_blog/models/ctype"
	"gvb_blog/response"
	"gvb_blog/service/user_service"
	"gvb_blog/utils"
	"gvb_blog/utils/jwt"
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
	if registerRes.Salt == "" {
		registerRes.Salt = userMo.Salt
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
	_claims, _ := ctx.Get("claims")
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
	err = global.DB.Model(&userMo).Update("password", PasswordRep.NewPassword).Error
	if err != nil {
		response.Fail(ctx, "修改失败")
		return
	}
	response.OkWithMessage(ctx, "修改成功")
}
