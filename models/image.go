package models

import (
	"gorm.io/gorm"
	"gvb_blog/global"
	"gvb_blog/models/ctype"
	"os"
)

type ImageModel struct {
	MODEL
	Path         string                 `json:"path"`                                  // 图片路径
	Hash         string                 `json:"hash"`                                  // 图片hash
	Name         string                 `gorm:"size:128" json:"name"`                  // 图片名称
	Suffix       string                 `gorm:"size:8" json:"suffix"`                  // 文件后缀
	Type         string                 `gorm:"size:8;default:'image'" json:"type"`    // 文件类型 image 或者 file
	FileLocation ctype.FileLocationType `gorm:"size:8;default:1" json:"file_location"` //默认是本地储存
}

// 删除钩子
func (I ImageModel) BeforeDelete(tx *gorm.DB) error {
	global.Log.Infof("正在删除图片:", I.FileLocation, ctype.Local)
	if I.FileLocation == ctype.Local {
		// 本地删除
		global.Log.Infof("删除的路径为:%v", I.Path)
		err := os.Remove(I.Path)
		if err != nil {
			global.Log.Error("删除数据失败:", err)
			return err
		}
	}
	return nil
}
