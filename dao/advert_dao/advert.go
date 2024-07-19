package advert_dao

import (
	"gvb_blog/global"
	"gvb_blog/models"
)

func DeleteUserList(ids []uint) (int64, error) {
	var am []*models.AdvertModel
	count := global.DB.Find(&am, ids).RowsAffected
	err := global.DB.Delete(&am).Error
	return count, err
}
