package models

// MenuImageModel 菜单图片关联表
type MenuImageModel struct {
	MenuID     uint       `json:"menu_id"`
	ImageID    uint       `json:"image_id"`
	MenuModel  MenuModel  `gorm:"foreignKey:MenuID;ON DELETE SET NULL" json:"menu_model"`
	ImageModel ImageModel `gorm:"foreignKey:ImageID;ON DELETE SET NULL" json:"image_model"`
	Sort       int        `gorm:"size:10" json:"sort"`
}
