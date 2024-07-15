package service

import "gvb_blog/models/ctype"

type ImageSortList struct {
	ImageId uint
	Sort    int
}

type MenuService struct {
	MenuTitle    string          `json:"menuTitle" binding:"required" msg:"完善标题"`        // 中文 =>导航条显示
	MenuTitleEn  string          `json:"menuTitleEn" binding:"required" msg:"完善标题的英文名称"` //英文=>路径显示
	Slogan       string          `json:"slogan"`                                         //
	Abstract     ctype.Array     `json:"abstract"`                                       //简介
	AbstractTime int             `json:"abstractTime"`                                   // 简介切换时间
	MenuTime     int             `json:"menuTime"`                                       //切换时间
	Sort         int             `json:"sort" binding:"required" msg:"完善菜单序号"`           //菜单序号
	ImageSort    []ImageSortList `json:"imageSort"`                                      // 图片排序顺序
}
