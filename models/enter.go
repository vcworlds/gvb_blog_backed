package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primary_key" json:"id"` // 主键ID
	CreatedAt time.Time `json:"created_at"`            // 创建时间
	UpdatedAt time.Time `json:"updated_at"`            // 更新时间
}

type Page struct {
	CurrentPage int    `form:"currentPage"` // 当前页
	Limit       int    `form:"limit"`       // 一页显示多少条数据
	Key         string `form:"key"`         // 搜索参数
	Sort        string `form:"sort"`        //排序
}
