package advert_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_blog/common"
	"gvb_blog/global"
	"gvb_blog/models"
	"gvb_blog/response"
	"gvb_blog/service/advert_service"
	"strings"
)

// Create
// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @Param data body service.AdvertResponse false "广告的参数"
// @Produce json
// @Router /advert/create  [post]
// @Success 200 {object} response.Response
func (a AdvertApi) Create(ctx *gin.Context) {
	var ar advert_service.AdvertResponse
	err := ctx.ShouldBindJSON(&ar)
	if err != nil {
		response.FailWithValidateError(err, &ar, ctx)
		return
	}
	res := ar.AdvertCreatService()
	if res.Code != 200 {
		response.Fail(ctx, res.Msg)
		return
	}
	response.OkWithMessage(ctx, res.Msg)
}

// Delete
// @Tags 广告管理
// @Summary 删除广告
// @Description 删除广告
// @Param data body  common.RemoveFileList true "删除广告所需参数"
// @Produce json
// @Router /advert/delete  [delete]
// @Success 200 {object} response.Response{data=string}
func (a AdvertApi) Delete(ctx *gin.Context) {
	var ids common.RemoveFileList
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		response.Fail(ctx, "数据绑定失败")
		return
	}
	res := advert_service.DeleteAdvertService(ids.Ids)
	if res.Code != 200 {
		response.Fail(ctx, res.Msg)
		return
	}
	response.OkWithMessage(ctx, fmt.Sprintf("删除成功,共删除了%d条数据", res.Data))
}

// Update
// @Tags 广告管理
// @Summary 编辑广告
// @Description 编辑广告
// @Param data body service.AdvertResponse true "编辑广告所需参数"
// @Produce json
// @Router /advert/update/:id  [put]
// @Success 200 {object} response.Response{data=string}
func (a AdvertApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var ads advert_service.AdvertResponse
	err := ctx.ShouldBindJSON(&ads)
	if err != nil {
		response.Fail(ctx, "数据绑定失败")
		return
	}
	res := ads.UpdateAdvertService(id)
	if res.Code != 200 {
		response.Fail(ctx, res.Msg)
		return
	}
	response.Ok(ctx, "更新成功", gin.H{
		"list": res.Data,
	})
}

// Show
// @Tags 广告管理
// @Summary 获取广告列表
// @Description 获取广告列表
// @Param data query models.Page true "获取广告列表的分页参数"
// @Produce json
// @Router /advert/show [get]
// @Success 200 {object} response.Response{data=string}
func (a AdvertApi) Show(ctx *gin.Context) {
	var advertPage models.Page
	err := ctx.ShouldBindQuery(&advertPage)
	if err != nil {
		global.Log.Error(err)
		response.Fail(ctx, "分页数据获取失败")
		return
	}
	referer := ctx.GetHeader("Referer")
	is_show := true
	if strings.Contains(referer, "admin") {
		is_show = false
	}
	list, count, _ := common.CommonPage(models.AdvertModel{IsShow: &is_show}, common.Option{advertPage})
	response.OkWithPage(ctx, list, count)

}
