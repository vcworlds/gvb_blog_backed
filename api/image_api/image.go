package image_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_blog/common"
	"gvb_blog/global"
	"gvb_blog/models"
	"gvb_blog/response"
	"gvb_blog/service/image_service"
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
	response.OkWithData(ctx, resList)
}

// Show 获取图片
func (ImageApi) Show(ctx *gin.Context) {
	var imagePage models.Page
	err := ctx.ShouldBindQuery(&imagePage)
	if err != nil {
		response.Fail(ctx, "分页数据绑定失败")
		return
	}
	list, count, err := common.CommonPage(models.ImageModel{}, common.Option{imagePage})
	if err != nil {
		response.Fail(ctx, "获取分页数据失败")
		return
	}
	response.OkWithData(ctx, gin.H{"count": count, "imageList": list})
}

// Delete 删除图片
func (a ImageApi) Delete(ctx *gin.Context) {
	var ids common.RemoveFileList
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		response.Fail(ctx, "绑定数据失败")
		return
	}
	var fileList []models.ImageModel
	count := global.DB.Find(&fileList, ids.Ids).RowsAffected
	if count == 0 {
		response.Fail(ctx, "没有找到图片信息")
		return
	}
	deleteRow := global.DB.Delete(&fileList).RowsAffected
	response.OkWithMessage(ctx, fmt.Sprintf("共删除了:%d条数据", deleteRow))
}

// Update 更新
func (a ImageApi) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
