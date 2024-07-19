package advert_service

type AdvertResponse struct {
	Title  string `json:"title"  binding:"required" msg:"请输入标题"`        // 广告标题 唯一
	Href   string `json:"href" binding:"required,url" msg:"广告链接不合规范"`   // 广告链接
	Images string `json:"images" binding:"required,url" msg:"图片链接不合规范"` // 图片
	IsShow *bool  `json:"is_show" binding:"required" msg:"请确定是否展示"`     // 是否显示
}
