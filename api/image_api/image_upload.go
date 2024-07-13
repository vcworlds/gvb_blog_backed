package image_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_blog/global"
	"gvb_blog/response"
	"path"
)

type ImageResponse struct {
	Name      string `json:"name"`
	IsSuccess bool   `json:"isSuccess"`
	Msg       string `json:"msg"`
}

func (ImageApi) ImageView(ctx *gin.Context) {
	fileForm, err := ctx.MultipartForm()
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	fileList, ok := fileForm.File["images"]
	if !ok {
		response.Fail(ctx, "文件参数获取失败")
		return
	}

	// 不存在就创建
	var resList []ImageResponse
	// 判断文件大小
	baseUrl := global.Config.Uploads.Path
	for _, file := range fileList {
		size := float64(file.Size) / float64(1024*1024)
		upPath := path.Join(baseUrl, file.Filename)
		if size > float64(global.Config.Uploads.Size) {
			resList = append(resList, ImageResponse{
				IsSuccess: false,
				Msg:       fmt.Sprintf("上传失败,超过文件上传最大限制:%dMB", global.Config.Uploads.Size),
				Name:      file.Filename,
			})
			continue
		}
		//保存图片
		err := ctx.SaveUploadedFile(file, upPath)
		if err != nil {
			resList = append(resList, ImageResponse{
				IsSuccess: false,
				Msg:       "上传失败",
				Name:      file.Filename,
			})
			continue
		}
		resList = append(resList, ImageResponse{
			IsSuccess: true,
			Msg:       "上传成功",
			Name:      upPath,
		})

	}
	response.OkWithData(ctx, resList)
}
