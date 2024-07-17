package user_service

type UserResponse struct {
	UserName string `json:"user_name" binding:"required" msg:"用户名不能为空"`
	Password string `json:"password" binding:"required" msg:"密码不能为空"`
	Token    string `json:"token"`
}

type RegisterRep struct {
	UserName   string `json:"user_name" binding:"required" msg:"用户名不能为空"`
	NickName   string `json:"nick_name" binding:"required" msg:"昵称不能为空"`
	Password   string `json:"password" binding:"required" msg:"密码不能为空"`
	RePassword string `json:"re_password" binding:"required" msg:"确认密码不能为空"`
	Email      string `json:"email" gorm:"type:email" binding:"required" msg:"邮箱不能为空"`
}
