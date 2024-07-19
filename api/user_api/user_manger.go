package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/common"
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

func (u UserApi) UserList(ctx *gin.Context) {
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

func (u UserApi) RoleUpdate(ctx *gin.Context) {

}
