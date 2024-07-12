package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/global"
	"gvb_blog/response"
)

func (SettingsApi) SettingsEmail(ctx *gin.Context) {
	response.OkWithData(ctx, global.Config.Email)
}

func (SettingsApi) SettingsEmailUpdate(ctx *gin.Context) {

}
