package dao

import (
	"gvb_blog/global"
)

func DeleteCommon[T any](model []T, ids []uint) (int64, error) {
	count := global.DB.Find(&model, ids).RowsAffected
	err := global.DB.Delete(&model).Error
	return count, err
}

func IsModelId[T any](model T, id string) (error, T) {
	err := global.DB.Take(&model, id).Error
	return err, model
}
