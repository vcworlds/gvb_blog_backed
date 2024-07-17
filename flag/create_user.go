package flag

import (
	"fmt"
	"gvb_blog/global"
	"gvb_blog/models"
	"gvb_blog/models/ctype"
	"gvb_blog/utils"
)

var (
	userName   string
	nickName   string
	Password   string
	RePassword string
)

func CreateUser(permissions string) {
	fmt.Println("请输入用户名：")
	fmt.Scan(&userName)
	fmt.Println("请输入昵称:")
	fmt.Scan(&nickName)
	fmt.Println("请输入密码:")
	fmt.Scan(&Password)
	fmt.Println("请确认密码:")
	fmt.Scan(&RePassword)
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		global.Log.Error("该用户已存在;请重新输入")
		return
	}
	if Password != RePassword {
		global.Log.Error("密码不正确;请重新输入")
		return
	}
	// 对密码进行加密
	hashPwd := utils.PasswordMd5(Password)
	// 判断角色
	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}
	// 入库
	global.DB.Create(&models.UserModel{
		MODEL:      models.MODEL{},
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		AvatarId:   1,
		Addr:       "内网注册",
		IP:         "127.0.0.1",
		Role:       role,
		SignStatus: ctype.SignEmail,
	})
}
