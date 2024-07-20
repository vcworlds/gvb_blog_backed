package advert_service

import (
	"gvb_blog/dao"
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
			Data: err,
		}
	}
	return response.Response{
		Code: http.StatusOK,
		Msg:  "创建成功",
		Data: nil,
	}
}

func DeleteAdvertService(ids []uint) *response.Response {
	count, err := dao.DeleteCommon[models.AdvertModel]([]models.AdvertModel{}, ids)
	res := &response.Response{
		Code: 200,
		Msg:  "删除成功",
		Data: nil,
	}
	if count == 0 {
		res.Code = http.StatusUnprocessableEntity
		res.Msg = "获取数据失败"
		return res
	}
	if err != nil {
		res.Code = http.StatusUnprocessableEntity
		res.Msg = "数据删除失败"
		return res
	}
	res.Data = count
	return res
}

func (a AdvertResponse) UpdateAdvertService(id string) *response.Response {
	res := &response.Response{
		Code: 200,
		Msg:  "",
		Data: nil,
	}
	// 判断id是否存在
	err, am := dao.IsModelId[models.AdvertModel](models.AdvertModel{}, id)
	if err != nil {
		res.Code = http.StatusUnprocessableEntity
		res.Msg = "查询数据失败"
		return res
	}
	err = global.DB.Model(&am).Updates(map[string]any{
		"title":   a.Title,
		"href":    a.Href,
		"images":  a.Images,
		"is_show": a.IsShow,
	}).Error
	if err != nil {
		res.Code = http.StatusUnprocessableEntity
		res.Msg = "更新数据失败"
		return res
	}
	res.Msg = "更新成功"
	res.Data = am
	return res
}
