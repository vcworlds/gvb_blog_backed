package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/global"
	"gvb_blog/response"
)

func (SettingsApi) SettingsInfo(ctx *gin.Context) {
	response.OkWithData(ctx, global.Config.SiteInfo)
}
