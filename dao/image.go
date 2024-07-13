package dao

import (
	"gvb_blog/global"
	"gvb_blog/models"
)

func ImageIsExit(imageHash string) (*models.ImageModel, error) {
	var image *models.ImageModel
	err := global.DB.Take(&image, "hash = ?", imageHash).Error
	if err == nil {
		//global.Log.Error(err)
		return image, err
	}
	return image, err
}
