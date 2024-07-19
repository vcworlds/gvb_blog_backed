package advert_service

import (
	"gvb_blog/global"
	"gvb_blog/models"
	"gvb_blog/response"
	"net/http"
)

func (a AdvertResponse) AdvertCreatService() response.Response {
	// 判断是否已经添加
	var am models.AdvertModel
	err := global.DB.Take(&am, "title = ?", a.Title).Error
	if err == nil {
		return response.Response{
			Code: http.StatusUnprocessableEntity,
			Msg:  "该条广告已添加",
			Data: nil,
		}
	}
	err = global.DB.Create(&models.AdvertModel{
		Title:  a.Title,
		Href:   a.Href,
		Images: a.Images,
		IsShow: a.IsShow,
	}).Error
	if err != nil {
		return response.Response{
			Code: http.StatusUnprocessableEntity,
			Msg:  "创建数据失败",
			Data: nil,
		}
	}
	return response.Response{
		Code: http.StatusOK,
		Msg:  "创建成功",
		Data: nil,
	}
}
