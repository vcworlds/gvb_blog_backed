package image_api

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/global"
	"gvb_blog/models"
	"gvb_blog/response"
	"gvb_blog/service"
)

func (ImageApi) ImageView(ctx *gin.Context) {
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
	resList := service.ImageService(fileList, ctx)
	response.OkWithData(ctx, resList)
}

func (ImageApi) ImageList(ctx *gin.Context) {
	var imagePage models.Page
	err := ctx.ShouldBindQuery(&imagePage)
	if err != nil {
		response.Fail(ctx, "分页数据绑定失败")
		return
	}
	var imageModel []models.ImageModel
	count := global.DB.Find(&imageModel).RowsAffected
	if imagePage.CurrentPage < 1 {
		imagePage.CurrentPage = 1
	}
	offset := (imagePage.CurrentPage - 1) * imagePage.Limit // 起始索引
	global.DB.Limit(imagePage.Limit).Offset(offset).Find(&imageModel)
	response.OkWithData(ctx, gin.H{"count": count, "imageList": imageModel})
}
