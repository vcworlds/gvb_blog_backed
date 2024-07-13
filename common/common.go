package common

import (
	"gvb_blog/global"
	"gvb_blog/models"
)

type Option struct {
	models.Page
}

func ComPage[T any](model T) (list []T, count int64) {
	count = global.DB.Find(&list).RowsAffected
	if
	global.DB.Limit().Offset().Find(&list)
}
