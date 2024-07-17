package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_blog/common"
	"gvb_blog/global"
	"gvb_blog/models"
	"gvb_blog/response"
	"gvb_blog/service"
)

// Create
// @Tags 菜单管理
// @Summary 创建菜单
// @Description 创建菜单
// @Param data body service.MenuService false "菜单的参数"
// @Produce json
// @Router /menu/create  [post]
// @Success 200 {object} response.Response
func (m MenuApi) Create(ctx *gin.Context) {
	var menuService service.MenuService
	err := ctx.ShouldBindJSON(&menuService)
	if err != nil {
		response.FailWithValidateError(err, &menuService, ctx)
		return
	}
	// 判断是否重复
	var menuM []models.MenuModel
	count := global.DB.Find(&menuM, "title = ? or path = ?", menuService.Title, menuService.Title).RowsAffected
	if count != 0 {
		response.Fail(ctx, "该路由已存在")
		return
	}
	// 创建表
	menuModel := &models.MenuModel{
		Title:        menuService.Title,
		Path:         menuService.Path,
		Slogan:       menuService.Slogan,
		Abstract:     menuService.Abstract,
		AbstractTime: menuService.AbstractTime,
		MenuTime:     menuService.MenuTime,
		Sort:         menuService.Sort,
	}
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		response.Fail(ctx, "菜单创建失败")
		return
	}
	if len(menuService.ImageSort) == 0 {
		response.Fail(ctx, "菜单排序有问题")
		return
	}
	// 创建关联表
	var menuImageList []models.MenuImageModel
	for _, sort := range menuService.ImageSort {
		menuImageList = append(menuImageList, models.MenuImageModel{
			MenuID:  menuModel.ID,
			ImageID: sort.ImageId,
			Sort:    sort.Sort,
		})
	}
	err = global.DB.Create(&menuImageList).Error
	if err != nil {
		response.Fail(ctx, "关联表失败")
		return
	}
	response.OkWithMessage(ctx, "创建成功")
}

// Delete
// @Tags 菜单管理
// @Summary 批量删除菜单
// @Description 批量删除菜单
// @Param data body common.RemoveFileList false "删除菜单的参数"
// @Produce json
// @Router /menu/delete  [delete]
// @Success 200 {object} response.Response
func (m MenuApi) Delete(ctx *gin.Context) {
	var ids common.RemoveFileList
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithValidateError(err, &ids.Ids, ctx)
		return
	}
	var menus []models.MenuModel
	count := global.DB.Find(&menus, "id in ?", ids.Ids).RowsAffected
	if count == 0 {
		response.OkWithMessage(ctx, "没找到相关数据")
		return
	}
	// 将对应的其他的模型清空
	err = global.DB.Model(&menus).Association("MenuImages").Clear()
	if err != nil {
		global.Log.Error(err)
		response.Fail(ctx, "删除关联表失败")
		return
	}
	deleteR := global.DB.Delete(&menus).RowsAffected
	response.OkWithMessage(ctx, fmt.Sprintf("共删除了了%d数据", deleteR))
}

// Update
// @Tags 菜单管理
// @Summary 编辑菜单
// @Description 编辑菜单
// @Param data body service.MenuService false "菜单的参数"
// @Produce json
// @Router /menu/update  [put]
// @Success 200 {object} response.Response
func (m MenuApi) Update(ctx *gin.Context) {
	menuId := ctx.Param("id")
	var menuRe service.MenuService
	var menuMo models.MenuModel
	err := global.DB.Take(&menuMo, menuId).Error
	if err != nil {
		response.Fail(ctx, "未查到该信息")
		return
	}
	err = ctx.ShouldBindJSON(&menuRe)
	if err != nil {
		response.FailWithValidateError(err, &menuRe, ctx)
		return
	}
	// 将关联表清空
	err = global.DB.Model(&menuMo).Association("MenuImages").Clear()
	if err != nil {
		response.Fail(ctx, "关联表更新失败")
		return
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
			response.Fail(ctx, "关联表更新失败")
			return
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
		response.Fail(ctx, "更新数据失败")
		return
	}
	response.OkWithMessage(ctx, "更新成功")

}

// Show
// @Tags 菜单管理
// @Summary 获取所有图片
// @Description 获取所有图片
// @Param data body string false "获取所有图片【不需要参数】"
// @Produce json
// @Router /menu/show  [get]
// @Success 200 {object} response.Response
func (m MenuApi) Show(ctx *gin.Context) {
	// 获取所有的菜单id
	var menuList []models.MenuModel
	var menuIds []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Order("sort desc").Scan(&menuIds)
	// 跟据菜单id查询关联的图片
	var menuImage []models.MenuImageModel
	global.DB.Preload("ImageModel").Order("sort desc").Find(&menuImage, "menu_id in ?", menuIds)
	var menuRes []service.MenuResponse
	// 循环每一个菜单
	for _, menu := range menuList {
		images := []service.Image{}
		// 循环关联表拿到每一个菜单所对应的图片
		for _, image := range menuImage {
			// 如果菜单id和关联表的菜单id不一样则退出循环
			if menu.ID != image.MenuID {
				continue
			}
			// 相同的话添加进去图片
			images = append(images, service.Image{
				Id:   image.ImageID,
				Path: image.ImageModel.Path,
			})
		}
		menuRes = append(menuRes, service.MenuResponse{
			MenuModel: menu,
			Image:     images,
		})
	}
	response.OkWithData(ctx, menuRes)

}

// MenuInfo
// @Tags 菜单管理
// @Summary 获取图片的基础信息id，path，title
// @Description 获取图片的基础信息
// @Param data body string false "获取图片的基础信息id【不需要参数】"
// @Produce json
// @Router /menu/menuInfo  [get]
// @Success 200 {object} response.Response
func MenuInfo(ctx *gin.Context) {
	var menuInfo []service.MenuInfo
	global.DB.Model(models.MenuModel{}).Select("id", "path", "title").Scan(&menuInfo)
	response.OkWithData(ctx, menuInfo)
}

func MenuDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	menuModel := models.MenuModel{}
	err := global.DB.Take(&menuModel, id).Error
	if err != nil {
		response.Fail(ctx, "查询详情失败")
		return
	}
	// 查询关联表
	var menuImage []models.MenuImageModel
	err = global.DB.Find(&menuImage, "menu_id = ?", id).Error
	if err != nil {
		response.Fail(ctx, "获取关联表失败")
		return
	}
	var imageList []service.Image
	for _, image := range menuImage {
		if menuModel.ID != image.MenuID {
			continue
		}
		imageList = append(imageList, service.Image{
			Id:   image.ImageID,
			Path: menuModel.Path,
		})
	}
	menuRe := &service.MenuResponse{
		MenuModel: menuModel,
		Image:     imageList,
	}
	response.OkWithData(ctx, menuRe)
}
