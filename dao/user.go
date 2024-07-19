package dao

import (
	"errors"
	"gvb_blog/global"
	"gvb_blog/models"
	"gvb_blog/utils"
)

// IsPassword 判断密码是否正确
func IsPassword(password string, userId uint) (bool, error) {
	var userMo models.UserModel
	err := global.DB.Take(&userMo, userId).Error
	if err != nil {
		return false, err
	}
	exits := utils.ValidPassword(password, userMo.Salt, userMo.Password)
	if !exits {
		return false, errors.New("密码不正确")
	}
	return true, nil
}
