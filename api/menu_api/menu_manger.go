package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/global"
	"gvb_blog/models"
	"gvb_blog/response"
	"gvb_blog/service"
)

type Image struct {
	Id   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Image `json:"banners"`
}

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
	// 创建表
	menuModel := &models.MenuModel{
		MenuTitle:    menuService.MenuTitle,
		MenuTitleEn:  menuService.MenuTitleEn,
		Slogan:       menuService.MenuTitle,
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

func (m MenuApi) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (m MenuApi) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

// Show
// @Tags 菜单管理
// @Summary 获取所有图片
// @Description 获取所有图片
// @Param data body  false "获取所有图片【不需要参数】"
// @Produce json
// @Router /menu/create  [post]
// @Success 200 {object} response.Response
func (m MenuApi) Show(ctx *gin.Context) {
	// 获取所有菜单中id
	var menuList []models.MenuModel
	var menuId []uint
	global.DB.Find(&menuList).Select("id").Order("sort desc").Scan(&menuId)
	var menuImage []models.MenuImageModel
	// 根据菜单id查找都里面所关联的图片
	global.DB.Preload("ImageModel").Order("sort desc").Find(&menuImage, "image_id in ?", menuId)
	// 返回数据
	var menus []MenuResponse
	// 循环获取每一个菜单
	for _, model := range menuList {
		// images
		var images []Image
		for _, image := range menuImage {
			if image.MenuID != model.ID {
				continue
			}
			images = append(images, Image{
				Id:   image.ImageID,
				Path: image.ImageModel.Path,
			})
			menus = append(menus, MenuResponse{
				MenuModel: model,
				Banners:   images,
			})
		}
		response.OkWithData(ctx, menus)
		return
	}

}
