package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/config"
	"gvb_blog/core"
	"gvb_blog/global"
	"gvb_blog/response"
)

func (SettingsApi) SettingsInfo(ctx *gin.Context) {
	response.OkWithData(ctx, global.Config.SiteInfo)
}

func (SettingsApi) SettingsUpdate(ctx *gin.Context) {
	var siteInfo config.SiteInfo
	err := ctx.ShouldBindJSON(&siteInfo)
	if err != nil {
		response.OkWithMessage(ctx, "参数错误")
		return
	}
	global.Config.SiteInfo = siteInfo
	err = core.SetYaml()
	if err != nil {
		global.Log.Fatalf("配置文件修改失败", err)
		return
	}
	response.OkWith(ctx)
}
