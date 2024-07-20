package menu_service

import (
	"gvb_blog/global"
	"gvb_blog/models"
	"gvb_blog/response"
	"net/http"
)

func (m MenuService) MenuCreateService() response.Response {
	// 判断是否重复
	var menuM []models.MenuModel
	count := global.DB.Find(&menuM, "title = ? or path = ?", m.Title, m.Title).RowsAffected
	res := response.Response{
		Code: http.StatusProcessing,
		Msg:  "",
		Data: nil,
	}
	if count != 0 {
		res.Msg = "该菜单已存在"
		return res
	}
	// 创建表
	menuModel := &models.MenuModel{
		Title:        m.Title,
		Path:         m.Path,
		Slogan:       m.Slogan,
		Abstract:     m.Abstract,
		AbstractTime: m.AbstractTime,
		MenuTime:     m.MenuTime,
		Sort:         m.Sort,
	}
	err := global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.Msg = "菜单创建失败"
		return res
	}
	if len(m.ImageSort) == 0 {
		res.Msg = "菜单排序有问题"
		return res
	}
	// 创建关联表
	var menuImageList []models.MenuImageModel
	for _, sort := range m.ImageSort {
		menuImageList = append(menuImageList, models.MenuImageModel{
			MenuID:  menuModel.ID,
			ImageID: sort.ImageId,
			Sort:    sort.Sort,
		})
	}
	err = global.DB.Create(&menuImageList).Error
	if err != nil {
		res.Msg = "关联表失败"
		return res
	}
	res.Code = http.StatusOK
	res.Msg = "创建成功"
	return res
}
func (menuRe MenuService) MenuUpdateService(menuMo models.MenuModel) response.Response {
	res := response.Response{
		Code: http.StatusProcessing,
		Msg:  "",
		Data: nil,
	}
	// 将关联表清空
	err := global.DB.Model(&menuMo).Association("MenuImages").Clear()
	if err != nil {
		global.Log.Error(err)
		res.Msg = "关联表更新失败"
		return res
	}
	// 创建关联表
	if len(menuRe.ImageSort) > 0 {
		var imageList []models.MenuImageModel
		for _, image := range menuRe.ImageSort {
			imageList = append(imageList, models.MenuImageModel{
				MenuID:  menuMo.ID,
				ImageID: image.ImageId,
				Sort:    image.Sort,
			})
		}
		err = global.DB.Create(&imageList).Error
		if err != nil {
			global.Log.Error(err)
			res.Msg = "关联表更新失败"
			return res
		}
	}
	// 普通更新
	err = global.DB.Model(&menuMo).Updates(map[string]any{
		"title":         menuRe.Title,
		"path":          menuRe.Path,
		"slogan":        menuRe.Slogan,
		"abstract":      menuRe.Abstract,
		"abstract_time": menuRe.AbstractTime,
		"menu_time":     menuRe.MenuTime,
		"sort":          menuRe.Sort,
	}).Error
	if err != nil {
		res.Msg = "更新数据失败"
		return res
	}
	res.Code = 200
	res.Msg = "更新成功"
	return res
}
