package models

import "gvb_blog/models/ctype"

type UserModel struct {
	MODEL
	NickName      string `json:"nickName" gorm:"size:32"`
	UserName      string `json:"userName" gorm:"size:32"`
	Password      string `json:"password" gorm:"size:128"`
	Salt          string `json:"salt"` // 密码盐
	Avatar        string `json:"avatar" gorm:"size:26"`
	Email         string `json:"email" gorm:"size:128"`
	Phone         string `json:"phone" gorm:"size:18"`
	Addr          string `json:"addr"`
	Role          ctype.Role
	Token         string           `json:"token" gorm:"size:64"`
	Ip            string           `json:"ip"`
	SignStatus    ctype.SignStatus `json:"signStatus"`
	ArticleModel  []ArticleModel   `json:"articleModel" gorm:"joinForeignKey:AuthId"`                                                    //文章列表
	CollectsModel []ArticleModel   `json:"collectsModel" gorm:"many2many:auth2_collects;joinForeignKey:AuthId;JoinReferences:ArticleId"` //收藏文章
}
