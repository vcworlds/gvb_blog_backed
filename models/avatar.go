package models

// 用户头像表
type Avatar struct {
	MODEL
	Path string `json:"path"`
}
