package image_service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_blog/dao"
	"gvb_blog/global"
	"gvb_blog/models"
	"gvb_blog/plugins/qiniu"
	"gvb_blog/utils"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

type ImageResponse struct {
	Name      string `json:"name"`
	IsSuccess bool   `json:"isSuccess"`
	Msg       string `json:"msg"`
}
type ImageListRes struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

func ImageService(fileList []*multipart.FileHeader, ctx *gin.Context) []ImageResponse {
	// 不存在就创建
	var resList []ImageResponse
	// 判断文件大小
	baseUrl := global.Config.Uploads.Path
	for _, file := range fileList {
		// 判断是否合法后缀
		pathExt := path.Ext(file.Filename)
		// 截取后缀
		suffix := strings.TrimPrefix(pathExt, ".")
		minSuffix := strings.ToLower(suffix)
		ok, _ := utils.InList(minSuffix, global.WhiteImageList)
		if !ok {
			resList = append(resList, ImageResponse{
				IsSuccess: false,
				Msg:       "不合规的图片后缀",
				Name:      file.Filename,
			})
			continue
		}
		size := float64(file.Size) / float64(1024*1024)
		// 组合文件路劲
		upPath := path.Join(baseUrl, file.Filename)
		if size > float64(global.Config.Uploads.Size) {
			resList = append(resList, ImageResponse{
				IsSuccess: false,
				Msg:       fmt.Sprintf("上传失败,超过文件上传最大限制:%dMB", global.Config.Uploads.Size),
				Name:      file.Filename,
			})
			continue
		}
		fileObj, err := file.Open()
		if err != nil {
			resList = append(resList, ImageResponse{Name: file.Filename, IsSuccess: false, Msg: "打开文件失败"})
			continue
		}
		fileData, err := io.ReadAll(fileObj)
		if err != nil {
			resList = append(resList, ImageResponse{Name: file.Filename, IsSuccess: false, Msg: "读取文件失败"})
			continue
		}
		imageHash := utils.Md5(fileData)
		// 判断图片内容是否存在数据库
		image, err := dao.ImageIsExit(imageHash)
		if err == nil {
			resList = append(resList, ImageResponse{
				IsSuccess: false,
				Msg:       fmt.Sprintf("该文件已存在,存在id为%d", image.ID),
				Name:      image.Path,
			})
			continue
		}
		// 判断是否将图片存在七牛存储
		if global.Config.QiNiu.Enable {
			filePath, err := qiniu.UploadImage(fileData, "gvb", file.Filename)
			if err != nil {
				global.Log.Error(err)
				continue
			}
			global.DB.Create(&models.ImageModel{
				Path:         filePath,
				Hash:         imageHash,
				Name:         file.Filename,
				Suffix:       minSuffix,
				FileLocation: 2,
			})
			resList = append(resList, ImageResponse{
				IsSuccess: true,
				Msg:       "上传七牛云成功",
				Name:      filePath,
			})
			continue
		}
		//保存图片
		err = ctx.SaveUploadedFile(file, upPath)
		if err != nil {
			resList = append(resList, ImageResponse{
				IsSuccess: false,
				Msg:       "上传失败",
				Name:      file.Filename,
			})
			continue
		}
		// 图片存入数据库
		global.DB.Create(&models.ImageModel{
			Path:         upPath,
			Hash:         imageHash,
			Name:         file.Filename,
			Suffix:       minSuffix,
			Type:         "本地",
			FileLocation: 1,
		})
		resList = append(resList, ImageResponse{
			IsSuccess: true,
			Msg:       "上传成功",
			Name:      upPath,
		})
	}
	return resList
}
