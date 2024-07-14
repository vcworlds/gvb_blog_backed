package advert_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_blog/common"
	"gvb_blog/global"
	"gvb_blog/models"
	"gvb_blog/response"
	"gvb_blog/service"
	"strings"
)

func (a AdvertApi) Create(ctx *gin.Context) {
	var ar service.AdvertResponse
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

func (a AdvertApi) Delete(ctx *gin.Context) {
	var ids common.RemoveFileList
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		response.Fail(ctx, "数据绑定失败")
		return
	}
	var am []models.AdvertModel
	count := global.DB.Find(&am, ids.Ids).RowsAffected
	if count == 0 {
		response.Fail(ctx, "获取数据失败")
		return
	}
	err = global.DB.Delete(&am).Error
	if err != nil {
		response.Fail(ctx, "数据删除失败")
		return
	}
	response.OkWithMessage(ctx, fmt.Sprintf("删除成功,共删除了%d条数据", count))
}

func (a AdvertApi) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	// 判断id是否存在
	var am models.AdvertModel
	err := global.DB.Take(&am, id).Error
	if err != nil {
		response.Fail(ctx, "该广告信息被删除了")
		return
	}
	var ads service.AdvertResponse
	err = ctx.ShouldBindJSON(&ads)
	if err != nil {
		response.Fail(ctx, "数据绑定失败")
		return
	}
	global.DB.Model(&am).Updates(map[string]any{
		"title":   ads.Title,
		"href":    ads.Href,
		"images":  ads.Images,
		"is_show": ads.IsShow,
	})
	response.Ok(ctx, "更新成功", gin.H{
		"list": am,
	})
}

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
