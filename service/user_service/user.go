package user_service

import (
	"context"
	"fmt"
	"gvb_blog/common"
	"gvb_blog/dao"
	"gvb_blog/global"
	"gvb_blog/models"
	"gvb_blog/models/ctype"
	"gvb_blog/response"
	"gvb_blog/utils"
	"gvb_blog/utils/jwt"
	"net/http"
	"time"
)

func (userRes LoginService) LoginService() response.Response {
	// 判断用户名是否存在
	var userMo models.UserModel
	err := global.DB.Take(&userMo, "user_name = ?", userRes.UserName).Error
	res := response.Response{
		Code: http.StatusFailedDependency,
		Msg:  "",
		Data: nil,
	}
	if err != nil {
		res.Msg = "未找到该用户"
		return res
	}
	// 判断密码是否正确
	exits, err := dao.IsPassword(userRes.Password, userMo.ID)
	if !exits || err != nil {
		global.Log.Error(err)
		res.Msg = "您的密码输入有误"
		return res
	}
	// 携带token
	token, err := jwt.ReleaseToken(userMo)
	if err != nil {
		global.Log.Error(err)
		res.Msg = "未知的错误,请刷新重试"
		return res
	}
	// 将token存储到redis
	expiration := global.Config.Jwt.Expires * 60 * 60
	id := fmt.Sprintf("token_%d", userMo.ID)
	if global.Redis == nil {
		global.Log.Error("global.Redis is nil")
		res.Msg = "Redis 初始化失败"
		return res
	}
	err = global.Redis.Set(context.Background(), id, token, time.Duration(expiration)*time.Second).Err()
	if err != nil {
		// 处理错误
		global.Log.Error("Error setting token in Redis:", err)
		res.Msg = "存储 Token 失败"
		return res
	}
	res.Code = http.StatusOK
	res.Msg = "登陆成功"
	res.Data = token
	return res
}

func (registerRes RegisterService) RegisterService() response.Response {
	res := response.Response{
		Code: http.StatusFailedDependency,
		Msg:  "",
		Data: nil,
	}
	// 判断用户名是否已经注册
	var userMo models.UserModel
	err := global.DB.Take(&userMo, "user_name = ?", registerRes.UserName).Error
	if err == nil {
		res.Msg = "该用户已存在"
		return res
	}
	// 判断两次密码是否一致
	if registerRes.RePassword != registerRes.Password {
		res.Msg = "两次密码不一致"
		return res
	}
	// 判断邮箱
	exits := utils.IsValidEmail(registerRes.Email)
	if !exits {
		res.Msg = "邮箱不合规范"
		return res
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
	res.Code = http.StatusOK
	res.Msg = "注册成功"
	return res
}

func UserListService(page models.Page, _claims any) response.Response {
	option := common.Option{page}
	userList, _, err := common.CommonPage[models.UserModel](models.UserModel{}, option)
	res := response.Response{
		Code: http.StatusFailedDependency,
		Msg:  "",
		Data: nil,
	}
	if err != nil {
		global.Log.Error(err)
		res.Msg = "获取列表失败"
		return res
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
	res.Code = 200
	res.Data = userModel
	return res
}

func (PasswordRep PasswordUpdateService) RegisterService(_claims any) response.Response {
	claims := _claims.(*jwt.Claims)
	var userMo models.UserModel
	err := global.DB.Take(&userMo, claims.UserId).Error
	res := response.Response{
		Code: http.StatusFailedDependency,
		Msg:  "",
		Data: nil,
	}
	if err != nil {
		res.Msg = "系统错误，请刷新重试"
		return res
	}
	// 判断密码是否正确
	exits, err := dao.IsPassword(PasswordRep.OldPassword, userMo.ID)
	if err != nil || !exits {
		res.Msg = "原始密码错误"
		return res
	}
	// 修改密码
	newPassword := utils.EncryptPassword(PasswordRep.NewPassword, userMo.Salt)
	err = global.DB.Model(&userMo).Update("password", newPassword).Error
	if err != nil {
		res.Msg = "修改密码错误，请刷新页面"
		return res
	}
	res.Code = 200
	res.Msg = "修改成功"
	return res
}
