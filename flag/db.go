package flag

import (
	"gvb_blog/global"
	"gvb_blog/models"
)

func MakeMigration() {
	var err error
	global.DB.SetupJoinTable(&models.UserModel{}, "collectsModels", &models.UserCollectsModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Image", &models.MenuImageModel{})

	global.Log.Infof("开始迁移数据库")
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.AdvertModel{},
		&models.ArticleModel{},
		&models.CommentModel{},
		&models.FeedbackModel{},
		&models.ImageModel{},
		&models.LoginDataModel{},
		&models.MenuModel{},
		&models.MenuImageModel{},
		&models.MessageModel{},
		&models.TagModel{},
		&models.UserModel{},
		&models.UserCollectsModel{},
	)
	if err != nil {
		global.Log.Error("数据库迁移失败: %v", err)
		return
	}
	global.Log.Info("数据库迁移成功")

}
