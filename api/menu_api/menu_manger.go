package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_blog/common"
	"gvb_blog/global"
	"gvb_blog/models"
	"gvb_blog/response"
	"gvb_blog/service/menu_service"
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
	var menuService menu_service.MenuService
	err := ctx.ShouldBindJSON(&menuService)
	if err != nil {
		response.FailWithValidateError(err, &menuService, ctx)
		return
	}
	res := menuService.MenuCreateService()
	if res.Code != 200 {
		response.Fail(ctx, res.Msg)
		return
	}
	response.OkWithMessage(ctx, res.Msg)
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
	var menuRe menu_service.MenuService
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
	res := menuRe.MenuUpdateService(menuMo)
	if res.Code != 200 {
		response.Fail(ctx, res.Msg)
	}
	response.OkWithMessage(ctx, res.Msg)

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
	var menuRes []menu_service.MenuResponse
	// 循环每一个菜单
	for _, menu := range menuList {
		images := []menu_service.Image{}
		// 循环关联表拿到每一个菜单所对应的图片
		for _, image := range menuImage {
			// 如果菜单id和关联表的菜单id不一样则退出循环
			if menu.ID != image.MenuID {
				continue
			}
			// 相同的话添加进去图片
			images = append(images, menu_service.Image{
				Id:   image.ImageID,
				Path: image.ImageModel.Path,
			})
		}
		menuRes = append(menuRes, menu_service.MenuResponse{
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
	var menuInfo []menu_service.MenuInfo
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
	var imageList []menu_service.Image
	for _, image := range menuImage {
		if menuModel.ID != image.MenuID {
			continue
		}
		imageList = append(imageList, menu_service.Image{
			Id:   image.ImageID,
			Path: menuModel.Path,
		})
	}
	menuRe := &menu_service.MenuResponse{
		MenuModel: menuModel,
		Image:     imageList,
	}
	response.OkWithData(ctx, menuRe)
}
