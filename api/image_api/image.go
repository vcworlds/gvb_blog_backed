package image_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_blog/common"
	"gvb_blog/dao/advert_dao"
	"gvb_blog/models"
	"gvb_blog/response"
	"gvb_blog/service/image_service"
	"net/http"
)

type IImageApi interface {
	response.RestApi
}
type ImageApi struct {
}

func NewImageApi() IImageApi {
	return ImageApi{}
}

// Create 添加图片
// @Tags 图片管理
// @Summary 添加图片
// @Description 添加图片
// @Param data query models.Page true "添加图片所需参数"
// @Produce json
// @Router /image [post]
// @Success 200 {object} response.Response{data=string}
func (ImageApi) Create(ctx *gin.Context) {
	fileForm, err := ctx.MultipartForm()
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	fileList, ok := fileForm.File["images"]
	//fmt.Println("x 的数据类型是:", reflect.TypeOf(fileList))
	if !ok {
		response.Fail(ctx, "文件参数获取失败")
		return
	}
	resList := image_service.ImageService(fileList, ctx)
	for _, res := range resList {
		if res.IsSuccess == false {
			response.Result(ctx, http.StatusFailedDependency, 422, res.Msg, res.Name)
			continue
		} else {
			response.Ok(ctx, res.Msg, map[string]any{
				"imageUrl": res.Name,
			})
		}
	}

}

// Show 获取图片
// @Tags 图片管理
// @Summary 获取图片列表
// @Description 获取图片列表
// @Param data query models.Page true "获取图片列表的分页参数"
// @Produce json
// @Router /imageList [get]
// @Success 200 {object} response.Response{data=string}
func (ImageApi) Show(ctx *gin.Context) {
	var imagePage models.Page
	err := ctx.ShouldBindQuery(&imagePage)
	if err != nil {
		response.Fail(ctx, "分页数据绑定失败")
		return
	}
	list, count, err := common.CommonPage[models.ImageModel](models.ImageModel{}, common.Option{imagePage})
	if err != nil {
		response.Fail(ctx, "获取分页数据失败")
		return
	}
	response.OkWithData(ctx, gin.H{"count": count, "imageList": list})
}

// Delete 删除图片
// @Tags 图片管理
// @Summary 批量删除
// @Description 批量删除
// @Param data query models.Page true "批量删除列表的参数"
// @Produce json
// @Router /imageList [delete]
// @Success 200 {object} response.Response{data=string}
func (a ImageApi) Delete(ctx *gin.Context) {
	var ids common.RemoveFileList
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		response.Fail(ctx, "绑定数据失败")
		return
	}
	count, err := advert_dao.DeleteUserList(ids.Ids)
	if count == 0 {
		response.Fail(ctx, "没有找到图片信息")
		return
	}
	if err != nil {
		response.Fail(ctx, "删除失败")
		return
	}
	response.OkWithMessage(ctx, fmt.Sprintf("共删除了:%d条数据", count))
}

// Update 更新
// @Tags 图片管理
// @Summary 更新图片
// @Description 获取广告列表
// @Param data query string true "更新"
// @Produce json
// @Router /advert/show [put]
// @Success 200 {object} response.Response{data=string}
func (a ImageApi) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
