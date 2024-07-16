package service

import (
	"gvb_blog/models"
	"gvb_blog/models/ctype"
)

type ImageSortList struct {
	ImageId uint
	Sort    int
}

type MenuService struct {
	Title        string          `json:"Title" binding:"required" msg:"完善标题"`     // 中文 =>导航条显示
	Path         string          `json:"Path" binding:"required" msg:"完善标题的英文名称"` //英文=>路径显示
	Slogan       string          `json:"slogan"`                                  //
	Abstract     ctype.Array     `json:"abstract"`                                //简介
	AbstractTime int             `json:"abstractTime"`                            // 简介切换时间
	MenuTime     int             `json:"menuTime"`                                //切换时间
	Sort         int             `json:"sort" binding:"required" msg:"完善菜单序号"`    //菜单序号
	ImageSort    []ImageSortList `json:"imageSort"`                               // 图片排序顺序
}

type Image struct {
	Id   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Image []Image `json:"images"`
}

type MenuInfo struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}
